package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() {
}

func cleanup() {
}

func TestExecutionExpectNoErrors(t *testing.T) {
	version := "latest"
	err := Execute(version)
	assert.Nil(t, err)
}
