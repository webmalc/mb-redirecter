package main

import (
	"testing"
	"webmalc/mb-redirector/common/test"

	"github.com/stretchr/testify/assert"
)

// Should return the config object.
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, 9000, c.Port)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
