package go_testing_test

import (
	"errors"
	"github.com/ldmarz/workshop-go-testing/go_testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_unnecessaryErrorHandler_shouldTestAllOutputsFromTableTest(t *testing.T) {
	type want struct {
		err    error
		result int
	}

	tests := []struct {
		name string
		a    int
		b    int
		want want
	}{
		{name: "should sum values when both are valid",
			a:    5,
			b:    5,
			want: want{result: 10},
		},
		{name: "should return error when a is negative",
			a:    -1,
			b:    5,
			want: want{err: errors.New("this amazing function only accepts positive numbers: -1 is negative")},
		},
		{name: "should return error when b is negative",
			a:    5,
			b:    -3,
			want: want{err: errors.New("this amazing function only accepts positive numbers: -3 is negative")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := go_testing.IncredibleSumFunction(tt.a, tt.b)

			if tt.want.err == nil {
				assert.Nil(t, err)
				assert.Equal(t, tt.want.result, result)
			} else {
				assert.EqualError(t, tt.want.err, err.Error())
			}
		})
	}
}
