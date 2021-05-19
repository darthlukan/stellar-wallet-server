package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	rw := httptest.NewRecorder()
	c, engine := gin.CreateTestContext(rw)

	engine.GET("/healthz", Health)
	request, err := http.NewRequestWithContext(c, http.MethodGet, "/healthz", nil)
	if err != nil {
		t.Fatalf("request: err = %v; want nil", err)
	}

	recorder := httptest.NewRecorder()
	engine.ServeHTTP(recorder, request)

	jdata := make(map[string]interface{})

	json.Unmarshal(recorder.Body.Bytes(), &jdata)

	assert.EqualValues(t, jdata["status"], http.StatusOK, "the two status should match")
}

/*
* Stub tests because we don't need to duplicate
 */

func TestPing(t *testing.T) {}

func TestCreateAccount(t *testing.T) {}

func TestGetAccount(t *testing.T) {}
