package main

import (
	"fmt"
)

type statusError struct {
    status string
}

func (stat statusError) Error() string {
    return fmt.Sprintf("status cannot be empty")
}

type numError struct {
	numb float64
}

func (num numError) Error() string {
	return fmt.Sprintf("status exceeds 140 characters")
}

// func validateStatus(status string) error {
// 	if (status == "") {
// 		return statusError{status: status}
// 	} else if (len(status) > 140) {
// 		return numError{numb: float64(len(status))}
// 	}
// 	return nil
// }

func validateStatus(status string) error {
	if (len(status) > 140) {
		return fmt.Errorf("status exceeds 140 characters")
	} else if (len(status) == 0) {
		return fmt.Errorf("status cannot be empty")
	}
	return nil
}