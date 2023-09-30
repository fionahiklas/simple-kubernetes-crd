package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/fionahiklas/simple-kubernetes-crd/electrictreesgo/internal/config"
	"github.com/fionahiklas/simple-kubernetes-crd/electrictreesgo/internal/treehandler"
	"github.com/fionahiklas/simple-kubernetes-crd/electrictreesgo/internal/version"
	"github.com/fionahiklas/simple-kubernetes-crd/electrictreesgo/pkg/logging"
	"github.com/fionahiklas/simple-kubernetes-crd/electrictreesgo/pkg/versionhandler"
)

func main() {
	parsedConfig, err := config.ParseConfig()
	if err != nil {
		fmt.Printf("Error parsing config: %s", err)
		os.Exit(1)
	}

	log := logging.NewLogger(parsedConfig)

	versionHandler := versionhandler.NewHandler(log, version.CodeVersion(), version.CommitHash())
	treeHandler := treehandler.NewHandler(log, parsedConfig)

	mux := http.NewServeMux()
	mux.Handle("/version", versionHandler)
	mux.Handle("/tree", treeHandler)

	log.Infof("Starting app version '%s', commit: '%s' on localhost:7777",
		version.CodeVersion(), version.CommitHash())

	if err := http.ListenAndServe("localhost:7777", mux); err != nil {
		log.Error(context.Background(), "Error from HTTP server: %s", err.Error())
	}
}
