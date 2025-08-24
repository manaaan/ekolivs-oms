resource "google_firestore_database" "default" {
  project     = var.project_id
  name        = "(default)"
  location_id = var.region
  type        = "FIRESTORE_NATIVE"
}

resource "google_firestore_index" "orders_creation_supplier" {
  project    = var.project_id
  database   = google_firestore_database.default.name
  collection = "orders"

  fields {
    field_path = "Supplier"
    order      = "ASCENDING"
  }

  fields {
    field_path = "CreationDate"
    order      = "DESCENDING"
  }

  fields {
    field_path = "__name__"
    order      = "DESCENDING"
  }
}

resource "google_firestore_index" "orders_creation_supplier" {
  project    = var.project_id
  database   = google_firestore_database.default.name
  collection = "demands"

  fields {
    field_path = "Member"
    order      = "ASCENDING"
  }

  fields {
    field_path = "CreationDate"
    order      = "DESCENDING"
  }

  fields {
    field_path = "__name__"
    order      = "DESCENDING"
  }
}
