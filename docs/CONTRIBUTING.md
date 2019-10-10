# Contributing


<!-- vim-markdown-toc GFM -->

* [Develop Locally](#develop-locally)
* [Testing](#testing)
  * [Unit Tests](#unit-tests)
  * [E2E Tests](#e2e-tests)
* [Create new API](#create-new-api)
* [Create new Controller for the API](#create-new-controller-for-the-api)
* [Create CRDs](#create-crds)
* [Create a new Release](#create-a-new-release)

<!-- vim-markdown-toc -->

## Develop Locally

Connect to a Kubernetes cluster

```shell
export KUBECONFIG=/path/to/config
```

Create all CRDs that are managed by the operator:

```shell
make init
```

Run Operator locally:

```shell
export ATLAS_PRIVATE_KEY=xxxx-xxxx-xxxx-xxxx
export ATLAS_PUBLIC_KEY=yyyyy
make
```

Create MongoDB Atlas Project

```shell
make deploy-project
```

Create MongoDB Atlas Cluster

```shell
make deploy-cluster
```

Delete MongoDB Atlas Project and Cluster

```shell
make delete-cluster
make delete-project
```

## Testing

### Unit Tests

The following executes unit tests for the controllers in `./pkg/controller/`

```shell
make test
# test only a subset
make test TEST_DIR=./pkg/controller/mongodbatlasdatabaseuser/...
# increase verbosity
make test TEST_DIR=./pkg/controller/mongodbatlasdatabaseuser/... VERBOSE="-v"
```

### E2E Tests

In order to run the end-to-end tests, you first have to create a namespace and a secret containing the private key of the programmatic API key pair which is needed by the Operator to perform API call against the MongoDB Atlas API.

The following command will execute the corresponding `kubectl` commands for you

```shell
export ATLAS_PRIVATE_KEY=xxxx-xxxx-xxxx-xxxx
make inite2etest
```

Afterwards, you can run the end-to-end tests with

```shell
export ATLAS_PUBLIC_KEY=yyyyy
make e2etest ORGANIZATION_ID=123456789
```

## Create new API

This example creates a new MongoDBAtlasCluster API:

```shell
make api KIND=MongoDBAtlasCluster
```

## Create new Controller for the API

To create a controller for the recently created API, run:

```shell
make controller KIND=MongoDBAtlasCluster
```

## Create CRDs

```shell
make generate-openapi
```

## Create a new Release

> You need to have Collaborator permissions to perform this step

A new release will

* create a new release on the Github [release page](https://github.com/Knappek/mongodbatlas-operator/releases) 
* push a new tagged Docker image to [Dockerhub](https://cloud.docker.com/repository/docker/knappek/mongodbatlas-operator/tags)

In order to do this, follow these steps:

1. Change the version in [.drone.yml](./.drone.yml) and in [operator.yaml](./deploy/operator.yaml) according to [Semantic Versioning](http://semver.org/)
2. Commit your changes (don't push)
3. Create a new release using [SemVer](http://semver.org/)

    ```shell
    make release VERSION=<major.minor.patch>
    ```

This will kick the CI pipeline and create a new Github Release with the version tag `v<major.minor.patch>`.

