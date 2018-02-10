package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/q231950/alethea/mocks"
	"github.com/q231950/alethea/model"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	server := Server{}
	assert.NotNil(t, server)
}

func TestPostStatusHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDataStorage := mocks.NewMockDataStorage(mockCtrl)
	server := NewServer(mockDataStorage)
	w := httptest.NewRecorder()
	err := errors.New("some error when creating the build result")
	incident := model.Incident{}
	server.handleBuildResult(incident, err, w)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)
}

func TestPostStatusHandlerRequiresBody(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDataStorage := mocks.NewMockDataStorage(mockCtrl)
	server := NewServer(mockDataStorage)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "http://example.com/foo", nil)
	server.postStatusHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusExpectationFailed, resp.StatusCode)
}

func TestPostStatusHandlerErrorsOnNonPostMethod(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDataStorage := mocks.NewMockDataStorage(mockCtrl)
	server := NewServer(mockDataStorage)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://example.com/foo", strings.NewReader("{\"json\":23}"))
	server.postStatusHandler(w, req)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusExpectationFailed)
}

func TestPostStatusHandlerCreatesStatusEntry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDataStorage := mocks.NewMockDataStorage(mockCtrl)
	mockDataStorage.EXPECT().
		StoreIncident(gomock.Any())

	server := NewServer(mockDataStorage)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "http://example.com/foo", strings.NewReader("{\"json\":23}"))
	server.postStatusHandler(w, req)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusAccepted)
}
