package calculator

import (
	"errors"
	"strconv"
)

// ErrDivisionByZero is returned when attempting to divide by zero
var ErrDivisionByZero = errors.New("division by zero")

// Add adds two float64 numbers
func Add(a, b float64) float64 {
	result := a + b

	return result
}

// Subtract subtracts b from a
func Subtract(a, b float64) float64 {
	result := a - b

	return result
}

// Multiply multiplies two float64 numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide divides a by b, returns an error if b is zero
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	return a / b, nil
}

// StringToFloat converts a string to float64
func StringToFloat(s string) (float64, error) {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

// FloatToString converts a float64 to string with specified precision
func FloatToString(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}
