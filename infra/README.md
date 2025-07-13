# Infrastructure-as-Code with OpenTofu

We decided to go with OpenTofu as fully Open Source alternative to terraform.

## Install tenv to manage OpenTofu version

https://github.com/tofuutils/tenv

```bash
tenv tofu install latest
```

### Apply infrastructure changes

To check for planned changes, run `tofu plan`.
To apply the changes, run `tofu apply`.

In the future the infra changes will be applied in CI/CD.

## Install gcloud

https://cloud.google.com/sdk/docs/install

### Authenticate

```bash
gcloud auth application-default login
```

Follow:
https://developer.hashicorp.com/terraform/tutorials/gcp-get-started/google-cloud-platform-build
