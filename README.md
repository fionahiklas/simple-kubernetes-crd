# Simple Kubernetes Custom Resource Definition

## Overview

Setup for testing Custom Resource Definition (CRD) and operator/controller.

The operator/CRD are for an object called "ElectricTrees".  Why?  Well,
my daughter loves the song [Together in Electric Dreams](https://en.wikipedia.org/wiki/Together_in_Electric_Dreams) but she thinks the line is 
"Together in Electric Trees" hence the name of this operator.


## Quickstart

* [Install Docker](./docs/install-docker.md#install)
* [Install Minikube](./docs/install-minikube.md#install)
* [Start Minikube](./docs/install-minikube.md#startup)
* [Build Electric Trees for Minikube](./electrictreesgo/README.md#quickstart)
* [Run the operator and apply sample CRD instance](./etoperator/README.md#quickstart)

You should now be able to run the following command to allow connections to 
the ElectricTrees service

```
minikube service --url=true electrictrees-mimi-svc
```

Running these commands to actually hit the endpoints

``` 
curl http://127.0.0.1:<service port>/tree  
curl http://127.0.0.1:<service port>/version  
```



## References

### Kubernetes

* [Operators intro video](https://www.youtube.com/watch?v=ha3LjlD6g7g)
* [Operator hub](https://operatorhub.io)
* [Operator SDK repo](https://github.com/operator-framework/operator-sdk)


