package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	assert.Equal(t, "test", os.Getenv("MB_REDIRECTOR_ENV"))
}

func TestMain(m *testing.M) {
	Run(m)
}
