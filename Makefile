.PHONY: build push deploy-staging deploy-prod migrate rollback delete-staging port-forward-staging deploy-staging-with-backup

DOCKER_REGISTRY ?= eduhass
VERSION ?= $(shell git rev-parse --short HEAD)

build:
	@docker buildx build --platform linux/amd64 -t ${DOCKER_REGISTRY}/kids-api:${VERSION} .
	@docker buildx build --platform linux/amd64 -t ${DOCKER_REGISTRY}/kids-api:latest .

push:
	@docker push ${DOCKER_REGISTRY}/kids-api:${VERSION}
	@docker push ${DOCKER_REGISTRY}/kids-api:latest

deploy-staging:
	@kustomize build kubernetes/overlays/staging | kubectl apply -f -

delete-staging:
	@kustomize build kubernetes/overlays/staging | kubectl delete -f -

port-forward-staging:
	@kubectl -n kids-app-staging port-forward svc/kids-api 8080:80

deploy-prod:
	@kubectl config use-context production
	@kustomize build kubernetes/overlays/production | kubectl apply -f -

migrate:
	@kubectl run kids-api-migrations --image=${DOCKER_REGISTRY}/kids-api:${VERSION} --restart=Never -- \
		/app/kids-api migrate up

rollback:
	@kubectl run kids-api-rollback --image=${DOCKER_REGISTRY}/kids-api:${VERSION} --restart=Never -- \
		/app/kids-api migrate down 1

deploy-staging-with-backup: deploy-staging
	@echo "Creating database backup..."
	@kubectl delete job postgres-backup -n kids-app-staging --ignore-not-found=true
	@kubectl apply -f kubernetes/base/postgres-backup-job.yaml -n kids-app-staging