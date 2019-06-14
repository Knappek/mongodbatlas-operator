BINARY = mongodbatlas-operator
COMMIT=$shell git rev-parse --short HEAD()
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE=$(shell date +%FT%T%z)

VERSION?=latest
OLM_VERSION?=0.1.0
API_VERSION?=v1alpha1
KIND=

DOCKERHUB_USERNAME=knappek

dev: generate-k8s
	operator-sdk up local
	
generate-k8s:
	operator-sdk generate k8s

api:
	operator-sdk add api --api-version=knappek.com/$(API_VERSION) --kind=$(KIND)

controller:
	operator-sdk add controller --api-version=knappek.com/$(API_VERSION) --kind=$(KIND)

docker-build: generate-k8s	
	operator-sdk build knappek/mongodbatlas-operator

docker-push: docker-build
	docker push $(DOCKERHUB_USERNAME)/$(BINARY):$(VERSION)

deploy-operator: docker-push
	kubectl delete deployment mongodbatlas-operator || true
	kubectl apply -f deploy/operator.yaml

project: cleanup
	kubectl create -f deploy/service_account.yaml
	kubectl create -f deploy/role.yaml
	kubectl create -f deploy/role_binding.yaml
	kubectl create -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_crd.yaml

deploy-project:
	kubectl apply -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml

delete-project:
	kubectl delete -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml

cleanup:
	kubectl delete mongodbatlasproject example-project >/dev/null 2>&1 || true
	kubectl delete -f deploy/ >/dev/null 2>&1 || true
	kubectl delete -f deploy/crds/ >/dev/null 2>&1 || true

olm-catalog:
	operator-sdk olm-catalog gen-csv --csv-version $(OLM_VERSION) --update-crds

test: cleanup olm-catalog
	operator-sdk scorecard \
		--cr-manifest deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml \
		--csv-path deploy/olm-catalog/mongodbatlas-operator/$(OLM_VERSION)/mongodbatlas-operator.v$(OLM_VERSION).clusterserviceversion.yaml