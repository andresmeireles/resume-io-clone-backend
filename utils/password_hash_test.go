package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordHash(t *testing.T) {
	t.Run("PasswordHashTest", func(t *testing.T) {
		cmd, err := PasswordHash("password")
		if err != nil {
			t.Error(err)
		}
		assert.NotEmpty(t, cmd)
	})

	t.Run("PasswordHashWithEmpty", func(t *testing.T) {
		_, err := PasswordHash("")
		assert.Equal(t, err.Error(), "Password cannot be empty")
	})
}
