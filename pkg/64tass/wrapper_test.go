package tass64_test

import (
	"testing"

	tass64 "github.com/devbay-io/unittest6502/pkg/64tass"
	"github.com/stretchr/testify/assert"
)

func TestWrapper(t *testing.T) {
	_, err := tass64.Run64tass("--help")
	assert.NoError(t, err, "should return help for 64tass")
	_, err = tass64.Run64tass("test")
	assert.Error(t, err, "64tass should throw error")
}

func Test64TassExists(t *testing.T) {
	tass64.TASS_EXECUTABLE = "64tass"
	assert.True(t, tass64.Does64TassExist(), "64tass command should exist")
	tass64.TASS_EXECUTABLE = "64tasse"
	assert.False(t, tass64.Does64TassExist(), "64tass command should not exist")
}
