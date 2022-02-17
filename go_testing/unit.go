package go_testing

import (
	"fmt"
)

func IncredibleSumFunction(a, b int) (int, error) {

	if err := unnecessaryErrorHandler(a, b); err != nil {
		return 0, fmt.Errorf("this amazing function only accepts positive numbers: %w", err)
	}

	return a + b, nil
}

func unnecessaryErrorHandler(a, b int) error {
	if a < 0 {
		return fmt.Errorf("%v is negative", a)
	}

	if b < 0 {
		// return errors.New("")
		return fmt.Errorf("%v is negative", b)
	}

	return nil
}
