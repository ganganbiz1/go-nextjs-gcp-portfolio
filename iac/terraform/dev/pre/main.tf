locals {
  environment = "dev"
  project = "template-pj-418409"
  region = "asia-northeast1"
  location = "asia-northeast1"
}

provider "google" {
  project     = local.project
  region      = local.region
}

module "artifact_registry_api" {
  source = "./../../modules/api/artifact_registry"
  location = local.location
}

module "artifact_registry_front" {
  source = "./../../modules/front/artifact_registry"
  location = local.location
}