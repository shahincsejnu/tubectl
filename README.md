# tubectl

This repository introduced `tubectl` which is basically similar to `kubectl`, `tubectl` offers several commands and features like `kubectl`. Currently supported commands of `tubectl` is given below.

[![Go Report Card](https://goreportcard.com/badge/github.com/shahincsejnu/client-go)](https://goreportcard.com/report/github.com/shahincsejnu/client-go)

## Environment

* Kind version: *kind v0.9.0 go1.15.2 linux/amd64*
* Kubernetes Server Version: *v1.19.1*

## How to use `tubectl`

* `git clone https://github.com/shahincsejnu/client-go.git`: clone this repo in your local machine
* `cd client-go`: go to the project client-go folder
* `go install`: create the binary of this project in your machine, which is named `tubectl`
* `tubectl`: It will print "Hello From tubectl", this is basically for confirming successful setup

## Deployment Commands

* `tubectl create deployment`: For creating a deployment (currently only default deployment will be created)
* `tubectl get deployments`: For getting all the deployments in the default namespace
* `tubectl delete deployment <deployment_name>`: For deleting the deployment named <deployment_name>
* `tubectl update deployment --deploy=<deployment_name> --image=<image_name> --replica=<replica_count>`: For updating the deployment named <deployment_name> with the <image_name> and <replica_count>

## Pod Commands

* `tubectl get pods`: For getting all the pods in default namespace
* `tubectl create pod`: For creating a pod (currently only default pod will be created)
* `tubectl delete pod <pod_name>`: For deleting <pod_name> pod
* `tubectl update pod --podname=<pod_name> --image=<image>`: For updating the pod named <pod_name> with the <image>