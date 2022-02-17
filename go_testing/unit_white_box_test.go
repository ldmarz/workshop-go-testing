package go_testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_unnecessaryErrorHandler_ShouldReturnErrorWhenAIsInvalid(t *testing.T) {
	err := unnecessaryErrorHandler(-1, 2)

	assert.EqualError(t, err, "-1 is negative")
}

func Test_unnecessaryErrorHandler_ShouldReturnErrorWhenBIsInvalid(t *testing.T) {
	err := unnecessaryErrorHandler(0, -5)

	assert.EqualError(t, err, "-5 is negative")
}
