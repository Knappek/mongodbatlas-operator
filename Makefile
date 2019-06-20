BINARY = mongodbatlas-operator
COMMIT=$shell git rev-parse --short HEAD()
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE=$(shell date +%FT%T%z)
CRDS=$(shell echo deploy/crds/*_crd.yaml | sed 's/ / -f /g')

VERSION?=latest
OLM_VERSION?=0.1.0
API_VERSION?=v1alpha1
KIND=

GITHUB_USERNAME=Knappek
DOCKERHUB_USERNAME=knappek

default: cleanup init dev

init: generate-openapi
	kubectl apply -f deploy/service_account.yaml
	kubectl apply -f deploy/role.yaml
	kubectl apply -f deploy/role_binding.yaml
	kubectl apply -f $(CRDS)

dev: generate-k8s
	operator-sdk up local
	
generate-k8s:
	operator-sdk generate k8s

generate-openapi:
	operator-sdk generate openapi

api:
	operator-sdk add api --api-version=knappek.com/$(API_VERSION) --kind=$(KIND)

controller:
	operator-sdk add controller --api-version=knappek.com/$(API_VERSION) --kind=$(KIND)

.PHONY: build
build:
	go build -o $(PWD)/build/_output/bin/$(BINARY) -gcflags all=-trimpath=${GOPATH} -asmflags all=-trimpath=${GOPATH} github.com/$(GITHUB_USERNAME)/$(BINARY)/cmd/manager

docker-build:
	docker build -f build/Dockerfile -t $(DOCKERHUB_USERNAME)/$(BINARY) .

docker-push: docker-build
	docker push $(DOCKERHUB_USERNAME)/$(BINARY):$(VERSION)

operator-build:
	operator-sdk build $(DOCKERHUB_USERNAME)/$(BINARY)

deploy-operator: docker-push
	kubectl delete deployment mongodbatlas-operator || true
	kubectl apply -f deploy/operator.yaml

deploy-project:
	kubectl apply -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml

delete-project:
	kubectl delete -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml

deploy-cluster:
	kubectl apply -f deploy/crds/knappek_v1alpha1_mongodbatlascluster_cr.yaml

delete-cluster:
	kubectl delete -f deploy/crds/knappek_v1alpha1_mongodbatlascluster_cr.yaml

cleanup:
	kubectl delete -f deploy/ >/dev/null 2>&1 || true
	kubectl delete -f deploy/crds/ >/dev/null 2>&1 || true

olm-catalog:
	operator-sdk olm-catalog gen-csv --csv-version $(OLM_VERSION) --update-crds

test: cleanup olm-catalog
	operator-sdk scorecard \
		--cr-manifest deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml \
		--csv-path deploy/olm-catalog/mongodbatlas-operator/$(OLM_VERSION)/mongodbatlas-operator.v$(OLM_VERSION).clusterserviceversion.yaml