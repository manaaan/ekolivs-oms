#!/bin/bash -e
# Ensure the list of API are enabled. Call with env variable PROJECT_ID set
# eg.
#   export PROJECT_ID=<...>
#   ./ensure-api "cloudfunctions.googleapis.com" "cloudbuild.googleapis.com"
#
PROJECT_ID=${PROJECT_ID?"Missing PROJECT_ID"}
APIS=( "$@" )

ENABLED_SERVICES_FILE=${ENABLED_SERVICES_FILE:-"$(mktemp)"}
if [ -s "$ENABLED_SERVICES_FILE" ]; then
  echo "Using cached file $ENABLED_SERVICES_FILE"
else
  gcloud services list \
    --enabled \
    --project="$PROJECT_ID" \
    > "$ENABLED_SERVICES_FILE"
fi

#Enable services
did_enabled_services="no"
for api in "${APIS[@]}"
do
  if grep -q "$api" "$ENABLED_SERVICES_FILE" ; then
    echo "Service '$api' is already enabled"
  else
    echo "Enable '$api' ..."
    did_enabled_services="yes"
    gcloud services enable "$api" --project="$PROJECT_ID"
  fi
done

if [[ "$did_enabled_services" = "yes" ]]; then
  echo "Waiting 30s after api was enabled to ensure it is activated"
  for _ in {1..10}; do
    echo -n '.'
    sleep 3
  done
  echo " Done"
fi
