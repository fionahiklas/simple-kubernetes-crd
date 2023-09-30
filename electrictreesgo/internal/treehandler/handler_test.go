package treehandler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fionahiklas/simple-kubernetes-crd/electrictreesgo/internal/treehandler"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {

	const (
		testTreeName   = "CountingPine"
		testTry        = true
		testHowFarAway = 42
		testEyesClosed = false
	)

	var mockResponseWriter *MockresponseWriter
	var mockLogger *Mocklogger
	var mockConfig *Mockconfig
	var testLogger *logrus.Logger

	resetTest := func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockResponseWriter = NewMockresponseWriter(ctrl)
		mockLogger = NewMocklogger(ctrl)
		mockConfig = NewMockconfig(ctrl)
		testLogger = logrus.New()
	}

	configHappyPath := func() {
		mockConfig.EXPECT().TreeName().Return(testTreeName)
		mockConfig.EXPECT().Try().Return(testTry)
		mockConfig.EXPECT().HowFarAway().Return(testHowFarAway)
		mockConfig.EXPECT().EyesClosed().Return(testEyesClosed)
	}

	t.Run("create new handler", func(t *testing.T) {
		resetTest(t)
		handler := treehandler.NewHandler(testLogger, mockConfig)
		require.NotNil(t, handler)
	})

	t.Run("bad request", func(t *testing.T) {
		resetTest(t)
		handler := treehandler.NewHandler(testLogger, mockConfig)
		testRequest := httptest.NewRequest("POST", "/wibble", nil)
		testResponseRecorder := httptest.NewRecorder()

		handler.ServeHTTP(testResponseRecorder, testRequest)
		require.Equal(t, http.StatusBadRequest, testResponseRecorder.Result().StatusCode)
	})

	t.Run("get request", func(t *testing.T) {
		resetTest(t)

		configHappyPath()

		handler := treehandler.NewHandler(testLogger, mockConfig)
		testRequest := httptest.NewRequest("GET", "/wibble", nil)
		testResponseRecorder := httptest.NewRecorder()

		handler.ServeHTTP(testResponseRecorder, testRequest)
		require.Equal(t, http.StatusOK, testResponseRecorder.Result().StatusCode)
	})

	t.Run("response write fails", func(t *testing.T) {
		resetTest(t)

		configHappyPath()

		handler := treehandler.NewHandler(mockLogger, mockConfig)
		testRequest := httptest.NewRequest("GET", "/wibble", nil)

		testError := errors.New("test error")
		mockResponseWriter.EXPECT().Header().Return(http.Header{})
		mockResponseWriter.EXPECT().WriteHeader(http.StatusOK)
		mockResponseWriter.EXPECT().Write(gomock.Any()).Return(0, testError)
		mockLogger.EXPECT().Errorf(gomock.Any(), testError).Times(1)
		mockLogger.EXPECT().Debugf(gomock.Any())

		handler.ServeHTTP(mockResponseWriter, testRequest)
	})

}
