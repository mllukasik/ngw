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
	err := Execute()
	assert.Nil(t, err)
}
