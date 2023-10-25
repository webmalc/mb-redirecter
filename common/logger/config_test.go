package logger

import (
	"testing"

	"webmalc/mb-redirector/common/test"

	"github.com/stretchr/testify/assert"
)

// Should return the config object.
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, true, c.IsDebug)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
