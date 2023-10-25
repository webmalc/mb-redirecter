package test

import (
	"os"
	"testing"

	"webmalc/mb-redirector/common/config"
)

// Setups the tests.
func setUp() {
	os.Setenv("MB_REDIRECTOR_ENV", "test")
	config.Setup()
}

// Run setups, runs and teardown the tests.
func Run(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}
