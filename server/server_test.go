package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/q231950/alethea/ci"
	"github.com/q231950/alethea/database/match"
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
	server := NewServer(mockDataStorage, "8080")
	w := httptest.NewRecorder()
	err := errors.New("some error when creating the build result")
	incident := &model.Incident{}
	server.handleBuildResult(incident, err, w)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)
}

func TestPostStatusHandlerRequiresBody(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDataStorage := mocks.NewMockDataStorage(mockCtrl)
	server := NewServer(mockDataStorage, "8080")
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "http://example.com/foo", nil)
	server.postStatusHandler(w, req, ci.Unknown)

	resp := w.Result()
	assert.Equal(t, http.StatusExpectationFailed, resp.StatusCode)
}

func TestPostStatusHandlerErrorsOnNonPostMethod(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDataStorage := mocks.NewMockDataStorage(mockCtrl)
	server := NewServer(mockDataStorage, "8080")
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://example.com/foo", strings.NewReader("{\"json\":23}"))
	server.postStatusHandler(w, req, ci.Unknown)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusExpectationFailed)
}

func TestPostStatusHandlerCreatesStatusEntry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDataStorage := mocks.NewMockDataStorage(mockCtrl)
	mockDataStorage.EXPECT().
		StoreCIBuild(gomock.Any())

	server := NewServer(mockDataStorage, "8080")
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "http://example.com/foo", strings.NewReader("{\"json\":23}"))
	server.postStatusHandler(w, req, ci.Circle)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusAccepted)
}

func TestPostCircleCIBuildStatusHandlerUsesCircleCIType(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDataStorage := mocks.NewMockDataStorage(mockCtrl)
	mockDataStorage.EXPECT().
		StoreCIBuild(match.CIType(ci.Circle))

	server := NewServer(mockDataStorage, "8080")
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "http://example.com/post/circle", strings.NewReader("{\"json\":23}"))
	server.postCircleCIBuildStatusHandler(w, req)
}
