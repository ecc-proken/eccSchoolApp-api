package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	signin := Signin{
		Status:  200,
		Message: "success",
	}

	assert.Equal(t, 200, signin.Status)
	assert.Equal(t, "success", signin.Message)
}
