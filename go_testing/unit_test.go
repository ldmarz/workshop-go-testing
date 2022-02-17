package go_testing_test

import (
	"github.com/ldmarz/workshop-go-testing/go_testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IncredibleSumFunction_ShouldReturnErrorWhenAIsNegative(t *testing.T) {
	_, err := go_testing.IncredibleSumFunction(-1, 0)
	if err == nil {
		panic("Should return an error when a parameter is negative")
	}
}

func Test_IncredibleSumFunction_Testify_ShouldReturnErrorWhenAIsNegative(t *testing.T) {
	_, err := go_testing.IncredibleSumFunction(-1, 0)

	assert.EqualError(t, err, "this amazing function only accepts positive numbers: -1 is negative")
}

func Test_IncredibleSumFunction_ShouldSumValuesWhenAreValid(t *testing.T) {
	result, err := go_testing.IncredibleSumFunction(2, 2)

	assert.Nil(t, err)
	assert.Equal(t, 4, result)
}

/*func Test_unnecessaryErrorHandler_ShouldReturnErrorWhenAIsInvalid(t *go_testing.T) {
	err := unnecessaryErrorHandler(-1, 2)

	assert.EqualError(t, err, "-1 is negative")
}*/
