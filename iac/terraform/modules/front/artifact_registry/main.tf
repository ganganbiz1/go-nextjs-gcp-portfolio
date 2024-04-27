resource "google_artifact_registry_repository" "front_repository" {
  location      = var.location
  repository_id = "dev-gcp-portfolio-repository-front"
  description   = "API Docker Image"
  format        = "DOCKER"

  labels = {
    env = "dev"
  }
}