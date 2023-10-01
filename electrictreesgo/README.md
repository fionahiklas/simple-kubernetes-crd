# Electric Trees Go

## Overview

Simple Go service that returns some data on an endpoint based on the config 
passed to it in the environment.


## Quickstart

* Ensure minikube is running
* Change to this directory 
* Run `make build_for_minikube`


## Development

* Run the following commands 

```
make install_tools
make generate
make test
make build
```

* Run locally with 

``` 
make run_local
```

* You can then hit the `/version` and `/tree` endpoints on port 7777 using curl


## References

### Go Code Snippets

* [simple-go-template](https://github.com/fionahiklas/simple-go-template)
* [simple-dynamodb-go-code](https://github.com/fionahiklas/simple-dynamodb-go-code)
* [simple-minikube-deployment](https://github.com/fionahiklas/simple-minikube-deployment)

### Docker

* [Golang base images](https://hub.docker.com/_/golang)


