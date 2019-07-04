# MongoDB Atlas Kubernetes Operator

[![Docker Pulls](https://img.shields.io/docker/pulls/knappek/mongodbatlas-operator.svg)](https://hub.docker.com/r/knappek/mongodbatlas-operator)
[![Build Status](https://cloud.drone.io/api/badges/Knappek/mongodbatlas-operator/status.svg)](https://cloud.drone.io/Knappek/mongodbatlas-operator)
[![Go Report Card](https://goreportcard.com/badge/github.com/Knappek/mongodbatlas-operator)](https://goreportcard.com/report/github.com/Knappek/mongodbatlas-operator)

## Overview

A Kubernetes Operator for [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) with which you can manage your MongoDB Atlas projects and clusters from within Kubernetes such as you do with your containerized applications. It is built using the [Operator Framework](https://github.com/operator-framework) and [Kubernetes Custom Resource Definitions (CRDs)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#customresourcedefinitions).

This project was inspired from the [MongoDB Atlas Terraform Provider](https://github.com/akshaykarle/terraform-provider-mongodbatlas) with the goal to have Kubernetes as the single source for both your (stateless) applications and MongoDB Atlas as the persistence layer. The benefit over using the Terraform provider is that `mongodbatlas-operator` ensures via Reconcile loops to have the desired state matching with the actual state and thus following the GitOps approach.

![](docs/mongodbatlas-operator-example.gif)


<!-- vim-markdown-toc GFM -->

* [Scope](#scope)
* [Prerequisites](#prerequisites)
* [Getting Started](#getting-started)
  * [Init](#init)
  * [Create a MongoDB Atlas Project](#create-a-mongodb-atlas-project)
  * [Create a Cluster](#create-a-cluster)
  * [List all MongoDB Atlas resources](#list-all-mongodb-atlas-resources)
* [Cleanup](#cleanup)
* [Developers Build Guide](#developers-build-guide)
* [Testing](#testing)
  * [E2E Tests](#e2e-tests)
* [Contributing](#contributing)
  * [Create new API](#create-new-api)
  * [Create new Controller for the API](#create-new-controller-for-the-api)
  * [Create CRDs](#create-crds)
  * [Create a new Release](#create-a-new-release)

<!-- vim-markdown-toc -->

## Scope

**Currently it supports**:

* Create/Delete MongoDB Atlas Projects
* Create/Delete MongoDB Atlas Clusters

## Prerequisites

* A running Kubernetes cluster, for example [Minikube](https://github.com/kubernetes/minikube) or [kind](https://github.com/kubernetes-sigs/kind)
* A MongoDB Atlas Account
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Getting Started

This example creates a MongoDB Atlas project and a cluster inside this project.

### Init

First, create the MongoDB Atlas project CRD and some RBAC:

```shell
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml
kubectl create -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_crd.yaml
kubectl create -f deploy/crds/knappek_v1alpha1_mongodbatlascluster_crd.yaml
```

Create a Kubernetes secret containing the Private Key of the [MongoDB Atlas Programmatic API Key](https://docs.atlas.mongodb.com/configure-api-access/#programmatic-api-keys)

```shell
kubectl create secret generic example-monogdb-atlas-project \
    --from-literal=privateKey=xxxxxxxxx
```

Adapt the environment variable `ATLAS_PUBLIC_KEY` in [operator.yaml](./deploy/operator.yaml) to your public key.

Deploy the MongoDB Atlas Project Operator:

```shell
kubectl apply -f deploy/operator.yaml
```

### Create a MongoDB Atlas Project

Adapt [knappek_v1alpha1_mongodbatlasproject_cr.yaml](./deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml) accordingly and deploy your first MongoDB Atlas Project

```shell
kubectl apply -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml
```

### Create a Cluster

Adapt [knappek_v1alpha1_mongodbatlascluster_cr.yaml](./deploy/crds/knappek_v1alpha1_mongodbatlascluster_cr.yaml) accordingly and deploy your first MongoDB Atlas Cluster

```shell
kubectl apply -f deploy/crds/knappek_v1alpha1_mongodbatlascluster_cr.yaml
```

### List all MongoDB Atlas resources

You can easily list all MongoDB Atlas related resources with

```shell
kubectl get mongodbatlas
```

## Cleanup

```shell
kubectl delete -f deploy/crds/knappek_v1alpha1_mongodbatlascluster_cr.yaml
kubectl delete -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml
kubectl delete -f deploy/
kubectl delete -f deploy/crds/
```

## Developers Build Guide

Connect to a Kubernetes cluster

```shell
export KUBECONFIG=/path/to/config
```

**Create all CRDs that are managed by the operator**

Run this once:

```shell
make init
```

**Run Operator locally**

```shell
make dev
```

**Create MongoDB Atlas Project**

```shell
make deploy-project
```

**Create MongoDB Atlas Cluster**

```shell
make deploy-cluster
```

**Delete MongoDB Atlas Project and Cluster**

```shell
make delete-cluster
make delete-project
```

## Testing

### E2E Tests

In order to run the end-to-end tests, you first have to create a namespace and a secret containing the private key of the programmatic API key pair which is needed by the Operator to perform API call against the MongoDB Atlas API.

The following command will execute the corresponding `kubectl` commands for you

```shell
export ATLAS_PRIVATE_KEY=xxxx-xxxx-xxxx-xxxx
export ATLAS_PUBLIC_KEY=yyyyy
make inite2etest
```

Afterwards, you can run the end-to-end tests with

```shell
make e2etest
```

## Contributing

Read through the [Contributing Guidelines and Code of Conduct](./CONTRIBUTING.md).

### Create new API

This example creates a new MongoDBAtlasCluster API:

```shell
make api KIND=MongoDBAtlasCluster
```

### Create new Controller for the API

To create a controller for the recently created API, run:

```shell
make controller KIND=MongoDBAtlasCluster
```

### Create CRDs

```shell
make generate-openapi
```

### Create a new Release

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
