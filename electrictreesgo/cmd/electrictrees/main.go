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

	versionHandler := versionhandler.NewHandler(log, version.AppName(), version.CodeVersion(), version.CommitHash())
	treeHandler := treehandler.NewHandler(log, parsedConfig)

	mux := http.NewServeMux()
	mux.Handle("/version", versionHandler)
	mux.Handle("/tree", treeHandler)

	log.Infof("Starting app version '%s', commit: '%s' on localhost:7777",
		version.CodeVersion(), version.CommitHash())

	// Note: Don't bind to localhost here because that can't be seen if this
	// app is running inside a container!  Localhost is a private network and
	// port mappings aren't going to work (well not unless we're dealing with
	// sidecar containers or other malarky)
	if err := http.ListenAndServe("0.0.0.0:7777", mux); err != nil {
		log.Error(context.Background(), "Error from HTTP server: %s", err.Error())
	}
}
