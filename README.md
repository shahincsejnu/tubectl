# tubectl

This repository introduced `tubectl` which is basically similar like `kubectl`, `tubectl` offers several commands and features like `kubectl`. Currently supported commands of `tubectl` is given below.

## Environment

* Kind version: 
* Kubernetes Server Version: 

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