# MongoDB Atlas Kubernetes Operator

[![Build Status](https://cloud.drone.io/api/badges/Knappek/mongodbatlas-operator/status.svg)](https://cloud.drone.io/Knappek/mongodbatlas-operator)
[![Go Report Card](https://goreportcard.com/badge/github.com/Knappek/mongodbatlas-operator)](https://goreportcard.com/report/github.com/Knappek/mongodbatlas-operator)
[![codecov](https://codecov.io/gh/Knappek/mongodbatlas-operator/branch/master/graph/badge.svg)](https://codecov.io/gh/Knappek/mongodbatlas-operator)

## Overview

A Kubernetes Operator for [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) with which you can manage your MongoDB Atlas projects and clusters from within Kubernetes such as you do with your containerized applications. It is built using the [Operator Framework](https://github.com/operator-framework) and [Kubernetes Custom Resource Definitions (CRDs)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#customresourcedefinitions).

This project was inspired from the [MongoDB Atlas Terraform Provider](https://github.com/akshaykarle/terraform-provider-mongodbatlas) with the goal to have Kubernetes as the single source for both your (stateless) applications and MongoDB Atlas as the persistence layer. The benefit over using the Terraform provider is that `mongodbatlas-operator` ensures via Reconcile loops to have the desired state matching with the actual state and thus following the GitOps approach.

![](docs/mongodbatlas-operator-example.gif)

<!-- vim-markdown-toc GFM -->

* [Scope](#scope)
* [Prerequisites](#prerequisites)
* [Getting Started](#getting-started)
  * [Deploy Operator](#deploy-operator)
  * [Create a MongoDB Atlas Project](#create-a-mongodb-atlas-project)
  * [Create a Cluster](#create-a-cluster)
  * [List all MongoDB Atlas resources](#list-all-mongodb-atlas-resources)
* [Cleanup](#cleanup)
* [Contributing](#contributing)

<!-- vim-markdown-toc -->

## Scope

**Currently it supports**:

* Create/Delete MongoDB Atlas Projects
* Create/Update/Delete MongoDB Atlas Clusters
* Create/Update/Delete MongoDB Atlas Database Users

## Prerequisites

* A running Kubernetes cluster, for example [Minikube](https://github.com/kubernetes/minikube) or [kind](https://github.com/kubernetes-sigs/kind)
* A MongoDB Atlas Account
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Getting Started

This example creates a MongoDB Atlas project and a cluster inside this project.

### Deploy Operator

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

## Environment Variables

You can specify the following environment variables in the Operator's [operator.yaml](./deploy/operator.yaml):


| Name | Description | Default | Required |
|------|-------------|---------|----------|
| WATCH_NAMESPACE | The namespace which the operator should watch for MongoDBAtlas CRDs. | `metadata.namespace` | yes |
| POD_NAME | Operator pod name. | `metadata.name` | no |
| OPERATOR_NAME | Operator name. | n/a | no |
| ATLAS_PRIVATE_KEY | The private key of the Atlas API. | n/a | yes |
| ATLAS_PUBLIC_KEY | The private key of the Atlas API. | n/a | yes |
| RECONCILIATION_TIME | Time in seconds which should be used to periodically reconcile the actual status in MongoDB Atlas with the current status in the corresponding Kubernetes CRD. | `"120"` | no |

## Contributing

I am working on this project in my spare time, hence feature development and release cycles could be improved ;). Contributors are welcome!

Read through the [Contributing Guidelines and Code of Conduct](./CONTRIBUTING.md).

More information how to contribute/develop can be found in the [docs](./docs/CONTRIBUTING.md).
