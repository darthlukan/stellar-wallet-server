package account

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Empty test because keypair.Random() doesn't require testing on our part.
func TestCreateKeyPair(t *testing.T) {}

func TestCreateAccount(t *testing.T) {
	rw := httptest.NewRecorder()
	c, engine := gin.CreateTestContext(rw)
	engine.POST("/accounts", CreateAccount)

	request, err := http.NewRequestWithContext(c, http.MethodPost, "/accounts", nil)
	if err != nil {
		t.Fatalf("request: err = %v; want nil", err)
	}

	recorder := httptest.NewRecorder()
	engine.ServeHTTP(recorder, request)

	jdata := make(map[string]interface{})

	json.Unmarshal(recorder.Body.Bytes(), &jdata)

	assert.EqualValues(t, jdata["status"], http.StatusOK, "the two status should match")
}
