package harness

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTestAccount(t *testing.T) {
	taccount, err := CreateTestAccount()
	if err != nil {
		t.Fatalf("CreateTestAccount():err = %v; want nil", err)
	}
	assert.Greater(t, len(taccount.Address), 0, "should be greater than zero")
	assert.Greater(t, len(taccount.Seed), 0, "should be reater than zero")
}

func TestFundTestAccount(t *testing.T) {
	taccount, err := CreateTestAccount()
	if err != nil {
		t.Fatalf("CreateTestAccount():err = %v; want nil", err)
	}

	rawAccountJson, err := FundTestAccount(&taccount)
	if err != nil {
		t.Fatalf("FundTestAccount(&taccount):err = %v; want nil", err)
	}

	var accountJson map[string]interface{}
	er := json.Unmarshal([]byte(rawAccountJson), &accountJson)
	if er != nil {
		t.Fatalf("json.Unmarshal:er = %v; want nil", er)
	}

	assert.True(t, accountJson["successful"].(bool), "successful should equal true")
	assert.Contains(t, accountJson["max_fee"].(string), "10", "max_fee should be greater than 10 XLM")
}
