package calculation

import (
	"errors"
)

// Division is for demonstration for go test
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}

	return a / b, nil
}
