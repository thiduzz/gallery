package models

import "strings"

const (
	ErrNotFound modelError = "models: resource not found"
	ErrInvalidID modelError = "models: user ID is invalid"

	ErrInvalidEmail modelError = "Email address is invalid"
	ErrRequiredEmail modelError = "Email address is required"
	ErrAlreadyTaken modelError = "models: email address is already taken"

	ErrRequiredPassword modelError = "Password is required"
	ErrInvalidPasswordHash modelError = "models: invalid password hash"
	ErrInvalidPassword modelError = "models: invalid password"

	ErrRememberTooShort modelError = "models: invalid remember token"
	ErrRequiredRememberHash modelError = "models: remember has in required"
)

type modelError string

func (m modelError) Error() string  {
	return string(m)
}

func (m modelError) Public() string  {
	s := strings.Replace(string(m), "models: ", "", 1)
	s = strings.Title(s[0:1]) + s[1:]
	return s
}