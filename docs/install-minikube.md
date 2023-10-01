# Install Minikube

## Overview

Steps for installing and configuring Minikube


## Install

### MacOS

Using [homebrew](https://brew.sh)

```
brew install minikube
```

Also added `kubectx` as this is useful to avoid long `kubectl` command that 
require the context and namespace being specified all the time

```
brew install kubectx
```

This installed the `kubectx` and `kubens` commands


### Ubuntu

Using Debian Packages

This is slightly more involved and requires downloading of packages before
installing them.

__NOTE:__ If you use `apt install <package file>` you could get an error like
this:

```
N: Download is performed unsandboxed as root as file '/home/fiona/kubectx_0.9.3-1_all.deb' couldn't be accessed by user '_apt'. - pkgAcquire::Run (13: Permission denied)
```

It seems to be harmless but can be avoided completely by using `dpkg` instead

Download the required packages using the following commands

```
wget http://ftp.uk.debian.org/debian/pool/main/k/kubectx/kubectx_0.9.3-1_all.deb
wget https://storage.googleapis.com/minikube/releases/latest/minikube_latest_arm64.deb
```

To install use the following commands

```
sudo dpkg -i ./kubectx_0.9.3-1_all.deb
sudo dpkg -i ./minikube_latest_arm64.deb
```

## Running Minikube

### Startup

Start minikube using the following command

```
minikube start --driver=docker
```


### Configuration

Enable the ingress addon

```
minikube addons enable ingress
```


## References

### Minikube

* [Download Page](https://minikube.sigs.k8s.io/docs/start/)

### Kubectx

* [Kubectx repo](https://github.com/ahmetb/kubectx)
* [Kubectx package page](https://packages.debian.org/bookworm/kubectx)
* [Debian Package](http://ftp.uk.debian.org/debian/pool/main/k/kubectx/kubectx_0.9.3-1_all.deb)

### Ubuntu

* [Unsandboxed error message](https://askubuntu.com/questions/908800/what-does-this-apt-error-message-download-is-performed-unsandboxed-as-root)

### Debian 

* [Package search](https://packages.debian.org/index)


