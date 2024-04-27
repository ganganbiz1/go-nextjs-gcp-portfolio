resource "google_cloud_run_v2_service" "api" {
  name     = "portfolio-service-api"
  location = var.location
  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    service_account = var.cloudrun_sa
    containers {
      # tagを指定しないとlatestが使われる
      image = "asia-northeast1-docker.pkg.dev/template-pj-418409/dev-gcp-portfolio-repository-api/portfolio-service-api"
      resources {
        limits = {
          cpu    = "1"
          memory = "512Mi"
        }
      }
      ports {
        container_port = 9000
      }
      env {
        name = "APP_ENV"
        value = "development"
      }
      env {
        name = "APP_NAME"
        value = "template-app-gcp"
      }
      env {
        name = "SERVER_PORT"
        value = "9000"
      }
      env {
        name = "POSTGRES_TIME_ZONE"
        value = "UTC"
      }
      env {
        name = "POSTGRES_SSL_MODE"
        value = "require"
      }
      env {
        name = "POSTGRES_HOST"
        value = "aws-0-ap-northeast-1.pooler.supabase.com"
      }
      env {
        name = "POSTGRES_DB"
        value = "postgres"
      }
      env {
        name = "POSTGRES_PORT"
        value = "5432"
      }
      env {
        name = "POSTGRES_USER"
        value = "postgres.sfxjcrcullssjqjkoxqy"
      }
      env {
        name = "POSTGRES_PASSWORD"
        value_source {
          secret_key_ref {
            secret = "API_DB_PASS"
            version = "latest"
          }
        }
      }
      env {
        name = "POSTGRES_HOST_REPLICA"
        value = "aws-0-ap-northeast-1.pooler.supabase.com"
      }
      env {
        name = "POSTGRES_DB_REPLICA"
        value = "postgres"
      }
      env {
        name = "POSTGRES_PORT_REPLICA"
        value = "5432"
      }
      env {
        name = "POSTGRES_USER_REPLICA"
        value = "postgres.sfxjcrcullssjqjkoxqy"
      }
      env {
        name = "POSTGRES_PASSWORD_REPLICA"
        value_source {
          secret_key_ref {
            secret = "API_DB_PASS"
            version = "latest"
          }
        }
      }
      env {
        name = "FIREBASE_CREDENTIALS_JSON"
        value_source {
          secret_key_ref {
            secret = "FIREBASE_CREDENTIALS_JSON"
            version = "latest"
          }
        }
      }
      env {
        name = "NEW_RELIC_LICENSE_KEY"
        value_source {
          secret_key_ref {
            secret = "NEW_RELIC_LICENSE_KEY"
            version = "latest"
          }
        }
      }
      env {
        name = "NEW_RELIC_ENABLED"
        value = "true"
      }
      env {
        name = "NEW_RELIC_TRACE_ENABLED"
        value = "true"
      }
      env {
        name = "NEW_RELIC_LOG_FORWARDING_ENABLED"
        value = "true"
      }
      env {
        name = "DD_AGENT_HOST"
        value = "datadog"
      }
      env {
        name = "DD_AGENT_PORT"
        value = "8126"
      }
    }
  }
}

resource "google_cloud_run_service_iam_binding" "public" {
  location    = google_cloud_run_v2_service.api.location
  project     = google_cloud_run_v2_service.api.project
  service     = google_cloud_run_v2_service.api.name
  role        = "roles/run.invoker"
  members     = ["allUsers"]
}