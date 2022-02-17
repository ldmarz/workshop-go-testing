package go_testing_test

import (
	"errors"
	"github.com/ldmarz/workshop-go-testing/go_testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSumClient struct {
	result int
	err    error
}

func (m mockSumClient) Sum(a, b int) (int, error) {
	if m.err != nil {
		return 0, m.err
	}

	return m.result, m.err
}

type TestifyMock struct {
	mock.Mock
}

func (t TestifyMock) Sum(a, b int) (int, error) {
	args := t.Called(a, b)

	err, ok := args.Get(0).(error)
	if ok {
		return 0, err
	}

	return args.Int(0), nil
}

func Test_SumWithExternalCall_ShouldReturnTheClientResponseWhenErrIsNil(t *testing.T) {
	sumClient := mockSumClient{
		result: 10,
		err:    nil,
	}

	service := go_testing.NewSumService(sumClient)

	result, err := service.SumWithExternalCall(5, 5)

	assert.Nil(t, err)
	assert.Equal(t, 10, result)
}

func Test_SumWithExternalCall_ShouldReturnErrorWhenClientFails(t *testing.T) {
	sumClient := mockSumClient{
		result: 10,
		err:    errors.New("unexpected error calling api"),
	}

	service := go_testing.NewSumService(sumClient)

	_, err := service.SumWithExternalCall(5, 5)

	assert.EqualError(t, err, "something happens calling external service: unexpected error calling api")
}

func Test_Testify_SumWithExternalCall_ShouldReturnErrorWhenClientFails(t *testing.T) {
	sumClient := new(TestifyMock)
	sumClient.On("Sum", 5, 5).Return(errors.New("unexpected error calling api"))
	service := go_testing.NewSumService(sumClient)

	_, err := service.SumWithExternalCall(5, 5)

	assert.EqualError(t, err, "something happens calling external service: unexpected error calling api")
}

func Test_Testify_SumWithExternalCall_ShouldReturnTheClientResponseWhenErrIsNil(t *testing.T) {
	sumClient := new(TestifyMock)
	sumClient.On("Sum", 5, 5).Return(10)
	service := go_testing.NewSumService(sumClient)

	result, err := service.SumWithExternalCall(5, 5)

	assert.Nil(t, err)
	assert.Equal(t, 10, result)
}
