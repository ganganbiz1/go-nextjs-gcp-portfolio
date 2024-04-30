ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
APP_COMPOSE_FILE := docker-compose.yml
APP_PROJECT_NAME := go-nextjs-gcp-portfolio-app
SWAGGER_COMPOSE_FILE := docker-compose_swagger.yml
SWAGGER_PROJECT_NAME := go-nextjs-gcp-portfolio-swagger
APP_CONTAINER_NAME := backend-api
SWAGGER_CONTAINER_NAME := swagger
TERRAFORM_CMD = docker run --rm -e GOOGLE_CREDENTIALS -v $(PWD)/iac/terraform:/terraform -w /terraform -it hashicorp/terraform:1.7.5 -chdir=/terraform/$(TF_ENV)/main
TERRAFORM_PRE_CMD = docker run --rm -e GOOGLE_CREDENTIALS -v $(PWD)/iac/terraform:/terraform -w /terraform -it hashicorp/terraform:1.7.5 -chdir=/terraform/$(TF_ENV)/pre

.PHONY: up
up: 
	docker-compose -f ${APP_COMPOSE_FILE} -p ${APP_PROJECT_NAME} up -d

.PHONY: up-build
up-build: 
	docker-compose -f ${APP_COMPOSE_FILE} -p ${APP_PROJECT_NAME} up -d --build

.PHONY: up-build-swagger
up-build-swagger: 
	docker-compose -f ${SWAGGER_COMPOSE_FILE} -p ${SWAGGER_PROJECT_NAME} up -d --build --remove-orphans

.PHONY: gen-oapi
gen-oapi: 
	rm -rf ./doc/swagger/openapi.yml
	rm -rf ./backend/openapi/openapi.yml
	docker-compose -f ${SWAGGER_COMPOSE_FILE} -p ${SWAGGER_PROJECT_NAME} up swagger-cli -d --build --remove-orphans
	@sleep 2 # コピーするまでに作成が間に合わないので、スリープ
	cp ./doc/swagger/openapi.yml ./backend/openapi/.

.PHONY: ui
ui: 
	docker-compose -f ${SWAGGER_COMPOSE_FILE} -p ${SWAGGER_PROJECT_NAME} up swagger-ui -d --build --remove-orphans

.PHONY: test
test: 
	docker container exec -it ${APP_CONTAINER_NAME} go test ./domain/service/... -coverprofile=coverage.out
	docker container exec -it ${APP_CONTAINER_NAME} go tool cover -html=coverage.out -o coverage.html

.PHONY: test-it
test-it: 
	docker container exec -it ${APP_CONTAINER_NAME} go test ./handler/...

.PHONY: wire
wire: 
	docker container exec -it ${APP_CONTAINER_NAME} go generate -x -tags wireinject ./wire/wire.go

.PHONY: wire-test
wire-test: 
	docker container exec -it ${APP_CONTAINER_NAME} go generate -x -tags wireinject ./wire/it/test_wire.go

.PHONY: oapi-codegen
oapi-codegen:
	docker container exec -it ${APP_CONTAINER_NAME} oapi-codegen -package openapi -generate types,server -o ./openapi/openapi.gen.go /src/backend/openapi/openapi.yml

.PHONY: lint
lint: 
	docker container exec -it ${APP_CONTAINER_NAME} golangci-lint run

.PHONY: mock-gen
mock-gen: 
	docker container exec -it ${APP_CONTAINER_NAME} sh -c 'cd mock && go generate ./...'

.PHONY: tf-build
tf-build:
	docker build -t tf-gcp -f ./iac/terraform/Dockerfile ./iac/terraform

.PHONY: init plan apply destroy

init:
	$(TERRAFORM_CMD) init

plan:
	$(TERRAFORM_CMD) plan

apply:
	$(TERRAFORM_CMD) apply

destroy:
	$(TERRAFORM_CMD) destroy

.PHONY: init-pre plan-pre apply-pre destroy-pre

init-pre:
	$(TERRAFORM_PRE_CMD) init

plan-pre:
	$(TERRAFORM_PRE_CMD) plan

apply-pre:
	$(TERRAFORM_PRE_CMD) apply

destroy-pre:
	$(TERRAFORM_PRE_CMD) destroy
