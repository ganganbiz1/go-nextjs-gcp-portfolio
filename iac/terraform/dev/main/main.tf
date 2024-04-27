locals {
  environment = "dev"
  project = "template-pj-418409"
  region = "asia-northeast1"
  location = "asia-northeast1"
}

provider "google" {
  // クレデンシャルは環境変数から読み込むのでコメントアウト
  #credentials = file("/terraform/dev/template-pj-418409-5a38afe49bb7.json")
  project     = local.project
  region      = local.region
}

module "iam" {
  source = "./../../modules/iam"
  project = local.project
}

module "cloudrun_api" {
  source = "./../../modules/api/cloudrun"
  location = local.location
  cloudrun_sa = module.iam.cloudrun_sa
}

module "cloudrun_front" {
  source = "./../../modules/front/cloudrun"
  location = local.location
  cloudrun_sa = module.iam.cloudrun_sa
}