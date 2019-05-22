BINARY = mongodbatlas-operator
COMMIT=$shell git rev-parse --short HEAD()
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE=$(shell date +%FT%T%z)

VERSION?=latest
DOCKERHUB_USERNAME=knappek

dev: generate-k8s
	operator-sdk up local
	
generate-k8s:
	operator-sdk generate k8s

docker-build: generate-k8s	
	operator-sdk build knappek/mongodbatlas-operator

docker-push: docker-build
	docker push $(DOCKERHUB_USERNAME)/$(BINARY):$(VERSION)

deploy: docker-push
	kubectl delete deployment mongodbatlas-operator || true
	kubectl apply -f deploy/operator.yaml

init-example-project:
	kubectl create -f deploy/service_account.yaml
	kubectl create -f deploy/role.yaml
	kubectl create -f deploy/role_binding.yaml
	kubectl create -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_crd.yaml

deploy-example-project:
	kubectl apply -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml

delete-example-project:
	kubectl delete -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml

cleanup:
	kubectl delete mongodbatlasproject example-project || true
	kubectl delete -f deploy/ || true
	kubectl delete -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_crd.yaml || trueg