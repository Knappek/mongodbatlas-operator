# MongoDB Atlas Kubernetes Operator

## Overview

A Kubernetes Operator for [MongoDB Atlas](LINK) with which you can manage your MongoDB Atlas projects and clusters from within Kubernetes such as you do with your containerized applications. It is built using the [Operator Framework](https://github.com/operator-framework) and [Kubernetes Custom Resource Definitions (CRDs)](LINK).
This project was inspired from the [MongoDB Atlas Terraform Provider](https://github.com/akshaykarle/terraform-provider-mongodbatlas) with the goal to have Kubernetes as the single source for both (stateless) applications and MongoDB Atlas as the persistence layer. Furthermore, the Kubernetes operator ensures via Reconcile loops to have the desired state matching with the actual state and thus following the GitOps approach. 

**Currently it supports**:

* Create MongoDB Atlas Project

**Coming soon**:

* Create/Update MongoDB Clusters
* Create/Update MongoDB Database Users
* Create/Update MongoDB Atlas Users
* Create/Update MongoDB Atlas Alerts
* Manage [Continuous Backups](LINK)

## Requirements

* A running Kubernetes cluster, for example Minikube
* A MongoDB Atlas Account (!!!LINK!!!)

## Usage

This example creates a MongoDB Atlas project:

First, create the MongoDB Atlas project CRD and some RBAC:

```shell
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml
kubectl create -f deploy/crds/mongodbatlas_v1alpha1_mongodbatlasproject_crd.yaml
```

Create a Kubernetes secret containing your MongoDB Atlas OrgID and the API Key

```shell
kubectl create secret generic example-monogdb-atlas-project \
    --from-literal=apiKey=xxxxxxxxx \
    --from-literal=orgId=yyyyyyyyyy
```

Deploy the MongoDB Atlas Project Operator:

```shell
kubectl apply -f deploy/operator.yaml
```

and finally deploy your first MongoDB Atlas Project

```shell
kubectl apply -f deploy/crds/knappek_v1alpha1_mongodbatlasproject_cr.yaml
```

## Developers Build Guide

### Initialize Operator

Run this once:

```shell
make init-example-project
```

### Run Operator locally

```shell
make dev
```

### Run Operator in Kubernetes Cluster

```shell
make deploy DOCKERHUB_USERNAME=<your-github-user>
```

## Contributing

TODO