terraform {
  backend "gcs" {
    bucket = "ekolivs-terraform-states"
    prefix = ""
  }
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">=5.31.1"
    }
  }
}

#Impersonate provider
provider "google" {
  alias   = "impersonate"
  project = var.project_id
  region  = var.region
  scopes = [
    "https://www.googleapis.com/auth/cloud-platform",
    "https://www.googleapis.com/auth/userinfo.email",
  ]
}

#Get token
data "google_service_account_access_token" "default" {
  provider               = google.impersonate
  target_service_account = var.service_account
  scopes                 = ["userinfo-email", "cloud-platform"]
  lifetime               = "3600s"
}

#Provider config
provider "google" {
  access_token = data.google_service_account_access_token.default.access_token
}
