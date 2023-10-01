# Electric Trees Operator

## Overview

A simple operator with a CRD to create an "electric trees" resource.

You don't have to carry out any of the steps in this documentation file,
this is what I did to get everything up and running.

__TODO:__ Write some unit tests for the operator!  The SDK seems to make the 
operator code eminently testable (defined input/output, injected client)

## Quickstart

* Run `cd etoperator` to get into this directory
* Run `make install` to install the CRD manifest
* Run `make deploy` NOTE: I don't think this step is actually needed
* Run `make run` this starts the operator locally

In a separate terminal run 

``` 
kustomize build config/samples | kubectl apply -f -
```

This will create the CRD instance and this should result in deployment and 
service being created, run `kubectl get all` to check, e.g.

``` 
kubectl get all                                     
NAME                                          READY   STATUS    RESTARTS   AGE
pod/electrictrees-mimi-app-67c4d579c7-7h7z4   1/1     Running   0          13m

NAME                             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
service/electrictrees-mimi-svc   NodePort    10.104.29.109   <none>        7777:32707/TCP   13m
service/kubernetes               ClusterIP   10.96.0.1       <none>        443/TCP          11h

NAME                                     READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/electrictrees-mimi-app   1/1     1            1           13m

NAME                                                DESIRED   CURRENT   READY   AGE
replicaset.apps/electrictrees-mimi-app-67c4d579c7   1         1         1       13m
```


## Environment

Most of this work was carried out on an M1 Mac with Docker Desktop, zsh, 
and homebrew for installing tools.  A few crucial points

* I have homebrew installed in `untar anywhere` mode because I don't like 
  tools having admin permissions or even needing them to install
* I use [direnv](https://direnv.net) to setup environment for this project 
* The GOPATH is set to `<repo>/.gopath` in a `.envrc` file that `direnv` picks up you need to `direnv allow .` at the root of the cloned repo for this to work


## Setup

### Install pre-requisites

#### MacOS

```
brew install direnv
brew install kustomize
```


### Install Operator SDK

#### MacOS

```
brew install operator-sdk
```


#### Build From Go

```
cd ~/wd/3rdparty

git clone https://github.com/operator-framework/operator-sdk
cd operator-sdk
git checkout master

# Ensure the install location is in the path 
make install
```


### Initialise

Ran these commands

```
cd etoperator

operator-sdk init --domain hiklas.com --owner "Fiona Bianchi" --repo github.com/fionahiklas/simple-kubernetes-crd/etoperator
```

__NOTE:__ If you have any backup files, for example `README.md~` the init command 
will complain 

```
Error: failed to initialize project: unable to run pre-scaffold tasks of "base.go.kubebuilder.io/v3": target directory is not empty (only go.mod, go.sum, files and directories with the prefix ".", files with the suffix ".md" or capitalized files name are allowed); found existing file "README.md~"
```

Also I got this warning

```
WARN[0000] the platform of this environment (darwin/arm64) is not suppported by kustomize v3 (v3.8.7) which is used in this scaffold. You will be unable to download a binary for the kustomize version supported and used by this plugin. The currently supported platforms are: ["linux/amd64" "linux/arm64" "darwin/amd64"] 
```

This is fixed by adding the `kustomize` prerequisite


### Creating API

Running this command

```
operator-sdk create api --version=v1alpha1 --kind=ElectricTrees
```

Creating the manifests

```
make manifests
```

### Coding

Following the Redhat operator tutorial for guidance on commands and code to creating a 
service and deployment from the ElectricTrees Custom Resource Definition


## References

### Operator SDK

* [Homepage](https://sdk.operatorframework.io)
* [Build operator in six steps](https://developers.redhat.com/articles/2021/09/07/build-kubernetes-operator-six-steps)
* [Redhat Operator SDK tutorial](https://developers.redhat.com/articles/2021/09/07/build-kubernetes-operator-six-steps)
* [Redhat Operator Tutorial code](https://github.com/deepak1725/hello-operator2)





