resource "google_artifact_registry_repository" "api_repository" {
  location      = var.location
  repository_id = "dev-gcp-portfolio-repository-api"
  description   = "API Docker Image"
  format        = "DOCKER"

  labels = {
    env = "dev"
  }
}