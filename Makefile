BINARY = mongodbatlas-operator
COMMIT=$(shell git rev-parse --short HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE=$(shell date +%FT%T%z)

VERSION?=latest
GITHUB_USERNAME=knappek

dev:
	operator-sdk up local
	
build: 
	operator-sdk build knappek/mongodbatlas-operator

docker-push: build
	docker push $(GITHUB_USERNAME)/$(BINARY):$(VERSION)

init-example-project:
	kubectl create -f deploy/service_account.yaml
	kubectl create -f deploy/role.yaml
	kubectl create -f deploy/role_binding.yaml
	kubectl create -f deploy/crds/mongodbatlas_v1alpha1_mongodbatlasproject_crd.yaml

deploy-example-project:
	kubectl apply -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml

delete-example-project:
	kubectl delete -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml