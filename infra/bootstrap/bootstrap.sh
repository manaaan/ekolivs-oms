#!/bin/bash
set -euo pipefail

# Setup Terraform in a gcp-project
PROGRAM="$0"
SCRIPTPATH="$( cd "$(dirname "$PROGRAM")" ; pwd -P )"

###############################################################################
NO_ARGS=0
E_OPTERROR=85

## Default values #############################################################

PROJECT_ID=ekolivs
GH_ORG=manaaan
REPO_NAME=ekolivs-oms
GOOGLE_USER=info@ekolivs.se

###############################################################################
USAGE="Usage: $PROGRAM -c <config-file> -p <project-id> -r <repo-name>

Setup terraform in a gcp-project

  -h               - Show this message
  -c <config-file> - path to config file
  -p <project-id>  - gcp project id
  -r <repo-name>   - name of the repo to connect the triggers to
  -g <group>       - Give access to all user in the group.

  example of config file examples/terraform-config.json
  the file should be checked in to the repo you are setting up

  Requires:
   - jq
   - gcloud
"

if [ $# -eq "$NO_ARGS" ]
then
  echo "$USAGE"
  exit "$E_OPTERROR"
fi

while getopts ":c:p:r:g:h" opt; do
  case $opt in
    c) CONFIG_FILE="$OPTARG"
    ;;
    p) PROJECT_ID="$OPTARG"
    ;;
    r) REPO_NAME="$OPTARG"
    ;;
    g) GOOGLE_GROUP="$OPTARG"
    ;;
    h) echo "$USAGE"
       exit 0
    ;;
    \?) echo "Invalid option -$OPTARG" >&2
    ;;
    *) echo "$USAGE" >&2 && exit 1;
  esac
done


########################################################################################
# UTILS
########################################################################################

RED='\033[0;31m'
NC='\033[0m' # No Color

print_warn() {
  printf "%s WARN: %s%s\n" "${RED}" "$1" "${NC}" >&2
}
print_err() {
  printf "%s ERROR: %s%s\n" "${RED}" "$1" "${NC}" >&2
  exit 1
}
_realpath() {
    [[ $1 = /* ]] && echo "$1" || echo "$PWD/${1#./}"
}

relpath() {
   python3 -c "import os.path; print( os.path.relpath('$1','${2:-$PWD}'))" ;
}

get_git_root() {
    (
        cd "$1"
        git rev-parse --show-toplevel
    )
}
get_remote_repo() {
    (
        cd "$1"
        # remove '.git' and 'git@github.com': from the url
        git remote get-url origin | sed -e 's/^git@github.com://' -e 's/\.git$//'
    )
}
export -f get_remote_repo
export -f get_git_root
long_env_name() {
  env="${1}"
  case "${1}" in
    dev|DEV|development) echo "development";;
    prod|PROD|production) echo "production";;
    *) echoerr "Invalid environment '${env}'. Only 'dev', 'development, 'prod' or 'production' allowed."
      exit 1;;
  esac
}


########################################################################################
# Variables
########################################################################################
CONFIG_FILE=$(_realpath "$CONFIG_FILE")
CONFIG_FILE_FOLDER="$(dirname "$CONFIG_FILE")"


export PROJECT_ID
#TODO use specific account instead of general permissive roles
ROLES_IMPERSONATE=(
   roles/iam.serviceAccountUser
   roles/iam.serviceAccountTokenCreator
)
PROJECT_LOCATION=EUROPE-NORTH1
GITHUB_SERVICE_ACCOUNT="github-infra@$PROJECT_ID.iam.gserviceaccount.com"
PROJECT_NUMBER="$(gcloud projects describe "${PROJECT_ID}" --format=json | jq -r .projectNumber)"

TERRAFORM_SERVICE_ACCOUNT="terraform@${PROJECT_ID}.iam.gserviceaccount.com"
TERRAFORM_STATES_BUCKET_NAME="${PROJECT_ID}-terraform-states"
TERRAFORM_STATES_BUCKET="gs://${TERRAFORM_STATES_BUCKET_NAME}"

########################################################################################
# VALIDATE input
########################################################################################
if [ -z "${CONFIG_FILE:-}" ]; then
  print_warn "Missing config file"
  echo "$USAGE" >&2 && exit 1;
fi

if [[ "$(get_remote_repo "$CONFIG_FILE_FOLDER")" != "$GH_ORG/$REPO_NAME" ]]; then
    print_err "Config file should be in the repo $REPO_NAME"
fi

########################################################################################

echo "Set project to $PROJECT_ID"
gcloud config set project "${PROJECT_ID}"

echo "enabling API:s"
APIS=$(jq -r '.apis[]' "$CONFIG_FILE")
# shellcheck disable=SC2086
"$SCRIPTPATH/ensure-api.sh" $APIS #skipping "" to allow for multiple arguments
echo "enabling API:s ... done"

#Create Terraform user
if gcloud iam service-accounts describe "${TERRAFORM_SERVICE_ACCOUNT}" > /dev/null 2>&1;
then
   echo "account already exists"
else
   echo "creating terraform service account"
   gcloud iam service-accounts create terraform --display-name "Terraform service account"
fi

#Add Roles
echo "adding roles"
if [ -z "${GOOGLE_GROUP:-}" ]; then
  member="user:${GOOGLE_USER}"
else
  member="group:${GOOGLE_GROUP}"
fi
roles=$(jq -r '.terraformSvcRoles[]' "$CONFIG_FILE")
# shellcheck disable=SC2086
"$SCRIPTPATH/ensure-roles.sh" "serviceAccount:${TERRAFORM_SERVICE_ACCOUNT}" $roles

current_terraform_roles=$(
   gcloud iam service-accounts get-iam-policy "${TERRAFORM_SERVICE_ACCOUNT}" \
     --project "${PROJECT_ID}" \
     --format json
)
echo "adding impersonate permissions"
for role in "${ROLES_IMPERSONATE[@]}"
 do
   if ! echo "$current_terraform_roles" | jq -e ".bindings[] | select(.role == \"$role\")" > /dev/null 2>&1; then
     echo "adding role $role"
      gcloud iam service-accounts add-iam-policy-binding "${TERRAFORM_SERVICE_ACCOUNT}" \
      --member "${member}" \
      --role "${role}"

   else
       echo "role $role already exists"
    fi
 done

# Create Bucket and block public access
if gcloud storage ls -b "$TERRAFORM_STATES_BUCKET" > /dev/null 2>&1 ;
then
   echo "bucket already exists"
else
   echo "creating bucket"
   gcloud storage buckets create \
      -l ${PROJECT_LOCATION} \
      --project "${PROJECT_ID}" \
      --pap "$TERRAFORM_STATES_BUCKET"
   #Enable versioning, Enable uniform access, Set bucket lifecycle policy
   gcloud storage buckets update "$TERRAFORM_STATES_BUCKET" \
      --uniform-bucket-level-access \
      --lifecycle-file="$SCRIPTPATH"/gcs-lifecycle-policy.json

   #Give access to the bucket
   gcloud storage buckets add-iam-policy-binding "$TERRAFORM_STATES_BUCKET" \
      --member "${member}" \
      --role=roles/storage.admin

fi

###############################################################################
# Create workload identity for github authentication
echo "Creating workload identiy pool"
if gcloud iam workload-identity-pools describe gh-pool --location=global --project "$PROJECT_ID" > /dev/null 2>&1 ; then
   echo "Pool already exists"
else
   echo "Creating pool"
   gcloud iam workload-identity-pools create "gh-pool" \
      --project="$PROJECT_ID" \
      --location="global" \
      --display-name="GH pool"
fi

echo "Creating workload identiy provider"
if gcloud iam workload-identity-pools providers describe gh-provider --location=global --workload-identity-pool=gh-pool --project "$PROJECT_ID" > /dev/null 2>&1 ; then
   echo "Provider already exists"
else
   echo "Creating provider"
   gcloud iam workload-identity-pools providers create-oidc "gh-provider" \
      --project="$PROJECT_ID" \
      --location="global" \
      --workload-identity-pool="gh-pool" \
      --display-name="GH provider" \
      --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.aud=assertion.aud,attribute.repository_owner=assertion.repository_owner,attribute.repository=assertion.repository" \
      --issuer-uri="https://token.actions.githubusercontent.com"
fi

echo "Creating github service account"
if gcloud iam service-accounts describe "$GITHUB_SERVICE_ACCOUNT" > /dev/null 2>&1; then
   echo "$GITHUB_SERVICE_ACCOUNT already exists"
else
   echo "Creating $GITHUB_SERVICE_ACCOUNT"
   gcloud iam service-accounts create "github-infra" \
      --project="$PROJECT_ID" \
      --description="Service account used by github actions" \
      --display-name="github-infra"
fi

echo "adding roles"
"$SCRIPTPATH/ensure-roles.sh" "serviceAccount:$GITHUB_SERVICE_ACCOUNT" \
   roles/iam.serviceAccountTokenCreator \
   roles/storage.admin \
   roles/secretmanager.secretAccessor \
   roles/viewer

gcloud iam service-accounts add-iam-policy-binding "$GITHUB_SERVICE_ACCOUNT" \
   --project="$PROJECT_ID" \
   --role="roles/iam.workloadIdentityUser" \
   --member="principalSet://iam.googleapis.com/projects/$PROJECT_NUMBER/locations/global/workloadIdentityPools/gh-pool/attribute.repository/${GH_ORG}/${REPO_NAME}"
