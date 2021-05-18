package account

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Empty test because keypair.Random() doesn't require testing on our part.
func TestCreateKeyPair(t *testing.T) {}

func TestCreateAccount(t *testing.T) {
	rw := httptest.NewRecorder()
	c, engine := gin.CreateTestContext(rw)
	engine.POST("/accounts", CreateAccount)

	var err error
	c.Request, err = http.NewRequest(http.MethodPost, "/accounts", nil)
	if err != nil {
		t.Fatalf("POST /accounts: err = %v; want nil", err)
	}

	engine.HandleContext(c)
}
