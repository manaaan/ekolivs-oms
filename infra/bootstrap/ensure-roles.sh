#!/bin/bash -e
# Ensure the list of roles are added to an user. Call with env variable PROJECT_ID set
# eg.
#   export PROJECT_ID=<...>
#   ./ensure-api <ACCOUNT> <ROLE_1> <ROLE_2>
#
set -euo pipefail
PROJECT_ID=${PROJECT_ID?"Missing PROJECT_ID"}
ACCOUNT=${1?"Missing ACCOUNT"}; shift 1
ROLES=( "$@" )

ENABLED_ROLES_FILE=${ENABLED_ROLES_FILE:-"$(mktemp)"}
if [ -s "$ENABLED_ROLES_FILE" ]; then
  echo "Using cached file $ENABLED_ROLES_FILE"
else
  gcloud projects get-iam-policy "$PROJECT_ID" \
    --format=json \
    > "$ENABLED_ROLES_FILE"
fi

#Enable services
for role in "${ROLES[@]}"
do
  if [ -s "$ENABLED_ROLES_FILE" ] && jq -ec ' .bindings[]
        | select( .role == "'"$role"'" )
        | .members[]
        |  select( . == "'"$ACCOUNT"'" )' "$ENABLED_ROLES_FILE" > /dev/null 2>&1;
  then
    echo "Role '$role' is already enabled for $ACCOUNT"
  else
    echo "Enable '$role' ..."
   gcloud projects add-iam-policy-binding "${PROJECT_ID}" \
        --member "${ACCOUNT}" \
        --role "${role}" \
        --condition=None
  fi

done
