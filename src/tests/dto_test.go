package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/xodiumx/go_back/api"
	"testing"
)

func TestUserResponseSerialization(t *testing.T) {
	// Create instance
	user := api.UserResponse{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// Serialize JSON
	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Error marshaling user response: %v", err)
	}

	// Check valid json
	expected := `{"id":1,"name":"John Doe","email":"john.doe@example.com"}`
	assert.JSONEq(t, expected, string(data))
}
