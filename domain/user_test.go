package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user := User{
		ID:       "2200000",
		Password: "hoge",
		UUID:		 "00000000-0000-0000-0000-000000000000",
	}
	
	assert.Equal(t, "2200000", user.ID)
	assert.Equal(t, "hoge", user.Password)
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", user.UUID)
}