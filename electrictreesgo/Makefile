
# Note: This Makefile can either be run from the command line/terminal or it
#       may be running inside a Dockerfile build.  In this second case some
#       of the environment variables may have been passed in as Docker build
#       arguments.

# Application name
APP_NAME=electrictrees

# Default to local target architecture if it's not being set by Docker
# build which will be calling this Makefile
TARGETARCH ?= $(shell uname -m)
TARGETOS ?= $(shell uname -s)

BUILD_PATH ?= $(CURDIR)/build
GO_MODULE_PATH ?= $(shell head -1 go.mod | cut -d' ' -f2-)

# This value may get passed in via an environment when we're building the
# Docker image otherwise we have to copy the whole .git folder and that
# is a bit of a waste of time
COMMIT_HASH ?= $(shell git rev-parse --short HEAD)
CODE_VERSION ?= $(shell cat ./VERSION)

GO_LD_FLAGS = -ldflags="-X ${GO_MODULE_PATH}/internal/version.commitHash=${COMMIT_HASH} -X ${GO_MODULE_PATH}/internal/version.codeVersion=${CODE_VERSION} -X ${GO_MODULE_PATH}/internal/version.appName=${APP_NAME}"

.PHONY: help tidy generate install_tools lint format test test_clean test_report test_summary build version

# from https://suva.sh/posts/well-documented-makefiles/
# sorting from https://stackoverflow.com/questions/14562423/is-there-a-way-to-ignore-header-lines-in-a-unix-sort
# Add double hash comments after every target to provide help text
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST) | awk 'NR<6{print $0; next}{print $0 | "sort"}'

version: ## Simply output the version number, used by some build scripts
	@echo "${CODE_VERSION}"

clean: ## Clean temporary/build files
	rm -f coverage.out
	find build -not -path '*/.*' -a -type f -print | xargs rm -f

tidy: ## Tidy imports
	go mod tidy

download: ## Ensure modules are downloaded
	go mod download
	go mod verify

install_tools: download ## Install tooling needed by other targets
	CGO_ENABLED=0 go install golang.org/x/tools/cmd/goimports
	CGO_ENABLED=0 go install github.com/golang/mock/mockgen
	# This tool actually uses the cgo support and shared libraries :/
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: ## Lint codebase
	golangci-lint -E goimports run internal/... pkg/... cmd/...

format: ## Format codebase and check imports
	goimports -w internal/ pkg/ cmd/

delete_mock: ## Deletes all files that match `mock_*.go`
	@echo "Deleting all mocks"
	find . -name "mock_*.go" -print -delete

generate: delete_mock ## Generate mocks
	@echo "Generating mocks ..."
	go generate ./internal/... ./pkg/...

clean_test_cache: ## Allows all tests to be forced to run without using cached results
	go clean -testcache

test_clean: clean_test_cache test ## Run tests but clear the cache first

test_report: test ## Run unit tests and generate an HTML coverage report
	go tool cover -html=coverage.out

test_summary: test ## Output coverage summary
	go tool cover -func=coverage.out

test: ## Run unit tests and check coverage
	@echo "Running unit tests"
	go test ${GO_LD_FLAGS} -coverprofile=coverage.out ./...

build: ## Build the application using native target or from inside Docker
	@echo "Building local '${APP_NAME}' for '${TARGETARCH}' on '${TARGETOS}'"
	CGO_ENABLED=0 go build ${GO_LD_FLAGS} -o ${BUILD_PATH}/${APP_NAME}-${TARGETOS}-${TARGETARCH} ./cmd/${APP_NAME}

run_local: build ## Run the application that has been built locally
	@echo "Running '${APP_NAME}' version '${CODE_VERSION}' locally on '${TARGETOS}/${TARGETARCH}'"
	( . ./scripts/local-env.sh ; ./build/${APP_NAME}-${TARGETOS}-${TARGETARCH} )

run_local_docker: ## Run the application in a Docker container that has been built locally
	@echo "Running Docker container for '${APP_NAME}' version '${CODE_VERSION}' locally on '${TARGETOS}/${TARGETARCH}'"
	( . ./scripts/local-env.sh ; docker run -d -p 7777:7777 --name ${APP_NAME} \
		--env LOG_LEVEL \
		--env ENABLE_JSON_LOG_FORMAT \
        --env TREE_NAME \
		--env TREE_TRY \
		--env TREE_DISTANCE \
		--env TREE_EYES_CLOSED \
		${APP_NAME}:${CODE_VERSION} )

kill_local_docker: ## Kill the running container and delete it
	@echo "Killing local docker container and removing it"
	docker kill ${APP_NAME}
	docker rm ${APP_NAME}

build_for_local: ## Build image using local docker daemon
	docker build --build-arg APP_NAME=${APP_NAME} --build-arg TARGETOS=${TARGETOS} --build-arg TARGETARCH=${TARGETARCH} --build-arg COMMIT_HASH=${COMMIT_HASH} -t ${APP_NAME}:${CODE_VERSION} .

build_for_minikube: ## Build the image using Minikubes docker daemon
	eval $$(minikube docker-env) ; docker build --build-arg APP_NAME=${APP_NAME} --build-arg TARGETOS=${TARGETOS} --build-arg TARGETARCH=${TARGETARCH} --build-arg COMMIT_HASH=${COMMIT_HASH} -t ${APP_NAME}:${CODE_VERSION} .
	eval $$(minikube docker-env) ; docker tag ${APP_NAME}:${CODE_VERSION} ${APP_NAME}:latest

check_minikube_images: ## Get a list of images from the Minikube docker daemon
	eval $$(minikube docker-env) ; docker images