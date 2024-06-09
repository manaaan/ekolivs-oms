## How to use setup terraform in a new project / repo | Usage of bootstrap.sh

### Requirements

Make sure the repository you're working on has been connected to Cloud Build. This process is manually done in GCP.

It requires some packages on your machine, only supports linux or macos

- `jq`
- `gcloud`

### Steps

1. Execute `bootstrap.sh` script

    ```bash
    ./bootstrap.sh \
      -c <path-to-config-in-repo>/terraform-config.json \
      -p <project> \
      -r <repo>
    ```
