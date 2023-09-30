//go:generate mockgen -package treehandler_test -destination=./mock_handler_test.go -source $GOFILE
package treehandler

import (
	"encoding/json"
	"net/http"
)

//lint:ignore U1000 Just here so that it can be mocked
type responseWriter interface {
	http.ResponseWriter
}

type logger interface {
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type config interface {
	TreeName() string
	Try() bool
	HowFarAway() int
	EyesClosed() bool
}

type handler struct {
	log    logger
	config config
}

func NewHandler(log logger, config config) *handler {
	return &handler{
		log:    log,
		config: config,
	}
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	h.log.Debugf("Tree handler called")
	if request.Method != http.MethodGet {
		h.log.Errorf("Tried to use method '%s'", request.Method)
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	responseJson := h.marshalTree()
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	_, err := responseWriter.Write(responseJson)

	if err != nil {
		h.log.Errorf("Couldn't write tree response, error: %w", err)
	}
}

// TODO: Cache the bytes so we don't keep marshalling the same thing
// TODO: I'd much prefer to marshal this as HTML with Lyrics but
// TODO: being a bit quicker for now
func (h *handler) marshalTree() []byte {
	// Don't forget, fields must be exported to be
	// marshalled (I forgot, hence this comment :) )
	type versionJson struct {
		TreeName   string `json:"tree_name,omitempty"`
		Try        bool   `json:"try,omitempty"`
		HowFarAway int    `json:"how_far,omitempty"`
		EyesClosed bool   `json:"eyes_closed,omitempty"`
	}

	toMarshal := versionJson{
		TreeName:   h.config.TreeName(),
		Try:        h.config.Try(),
		HowFarAway: h.config.HowFarAway(),
		EyesClosed: h.config.EyesClosed(),
	}

	// I can't find a way that this can actually fail given the input
	// is strings.  The tests should handle catching anything really bizarre
	// or the app will panic fairly quickly when running is the version
	// endpoint is used as a health check in Docker/k8s
	marshalledVersion, _ := json.Marshal(&toMarshal)
	return marshalledVersion
}
