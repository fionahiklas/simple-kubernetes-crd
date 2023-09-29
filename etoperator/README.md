# Electric Trees Operator

## Overview

A simple operator with a CRD to create an "electric trees" resource.

You don't have to carry out any of the steps in this documentation file,
this is what I did to get everything up and running.


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

NOTE: If you have any backup files, for example `README.md~` the init command 
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





## References

### Operator SDK

* [Homepage](https://sdk.operatorframework.io)
* [Build operator in six steps](https://developers.redhat.com/articles/2021/09/07/build-kubernetes-operator-six-steps)








