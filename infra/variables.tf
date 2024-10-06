variable "project_id" {
  description = "Project ID"
  type        = string
}

variable "project_number" {
  description = "Project number"
  type        = string
}

variable "region" {
  description = "Default region"
  type        = string
}

variable "service_account" {
  description = "Service account to impersonate"
  type        = string
}

variable "gh_org" {
  description = "GitHub organization"
  type        = string

}

variable "gh_repo" {
  description = "GitHub repository"
  type        = string

}
