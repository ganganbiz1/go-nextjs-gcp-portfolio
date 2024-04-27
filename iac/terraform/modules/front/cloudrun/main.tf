resource "google_cloud_run_v2_service" "front" {
  name     = "portfolio-service-front"
  location = var.location
  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    service_account = var.cloudrun_sa
    containers {
      # tagを指定しないとlatestが使われる
      image = "asia-northeast1-docker.pkg.dev/template-pj-418409/dev-gcp-portfolio-repository-front/portfolio-service-front"
      resources {
        limits = {
          cpu    = "1"
          memory = "512Mi"
        }
      }
      ports {
        container_port = 3000
      }
    }
  }
}

resource "google_cloud_run_service_iam_binding" "public" {
  location    = google_cloud_run_v2_service.front.location
  project     = google_cloud_run_v2_service.front.project
  service     = google_cloud_run_v2_service.front.name
  role        = "roles/run.invoker"
  members     = ["allUsers"]
}