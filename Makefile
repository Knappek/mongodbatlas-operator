BINARY = mongodbatlas-operator
COMMIT=$shell git rev-parse --short HEAD()
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE=$(shell date +%FT%T%z)
CRDS=$(shell echo deploy/crds/*_crd.yaml | sed 's/ / -f /g')
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GO := GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go

ORGANIZATION_ID?=5c4a2a55553855344780cf5f

VERSION?=latest
OLM_VERSION?=0.0.4
API_VERSION?=v1alpha1
KIND=

GITHUB_USERNAME=Knappek
DOCKERHUB_USERNAME=knappek

default: cleanup init dev

init:
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
	$(GO) build -o $(PWD)/build/_output/bin/$(BINARY) -gcflags all=-trimpath=${GOPATH} -asmflags all=-trimpath=${GOPATH} github.com/$(GITHUB_USERNAME)/$(BINARY)/cmd/manager

docker-build:
	docker build -f build/Dockerfile -t $(DOCKERHUB_USERNAME)/$(BINARY) .

docker-push: docker-build
	docker push $(DOCKERHUB_USERNAME)/$(BINARY):$(VERSION)

release:
	git tag v${VERSION}
	git push && git push --tags

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

csv:
	operator-sdk olm-catalog gen-csv --csv-version $(OLM_VERSION) --update-crds

e2etest: cleanup
	operator-sdk test local ./test/e2e \
		--namespace default \
		--go-test-flags "-v --organizationID=$(ORGANIZATION_ID)"

fmt:
	gofmt -w $(GOFMT_FILES)

lint:
	@which golint > /dev/null; if [ $$? -ne 0 ]; then \
		$(GO) get -u golang.org/x/lint/golint; \
	fi
	 $(GO) list ./... | grep -v /vendor/ | xargs golint -set_exit_status