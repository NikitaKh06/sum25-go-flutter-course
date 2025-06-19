package user

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrInvalidEmail is returned when the email format is invalid
	ErrInvalidEmail = errors.New("invalid email format")
	// ErrInvalidAge is returned when the age is invalid
	ErrInvalidAge = errors.New("invalid age: must be between 0 and 150")
	// ErrEmptyName is returned when the name is empty
	ErrEmptyName = errors.New("name cannot be empty")
)

// User represents a user in the system
type User struct {
	Name  string
	Age   int
	Email string
}

// NewUser creates a new user with validation
func NewUser(name string, age int, email string) (*User, error) {

	newUser := &User{
		Name:  name,
		Age:   age,
		Email: email,
	}

	err := newUser.Validate()
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrEmptyName
	}

	if u.Age < 0 || u.Age > 150 {
		return ErrInvalidAge
	}

	if !IsValidEmail(u.Email) {
		return ErrInvalidEmail
	}

	return nil
}

// String returns a string representation of the user
func (u *User) String() string {
	stringRepresentation := fmt.Sprintf("Name: %s, Age: %d, Email: %s", u.Name, u.Age, u.Email)

	return stringRepresentation
}

// IsValidEmail checks if the email format is valid
func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}

	at := strings.Index(email, "@")
	if at <= 0 {
		return false
	}

	dot := strings.LastIndex(email, ".")
	if dot <= at+1 {
		return false
	}

	if dot >= len(email)-2 {
		return false
	}

	return true
}
