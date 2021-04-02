package models

import "strings"

const (
	ErrNotFound modelError = "models: resource not found"
	ErrInvalidEmail modelError = "Email address is invalid"
	ErrRequiredEmail modelError = "Email address is required"
	ErrAlreadyTaken modelError = "models: email address is already taken"
	ErrRequiredPassword modelError = "Password is required"
	ErrInvalidPassword modelError = "models: invalid password"
	ErrRequiredTitle modelError = "models: title is required"

	ErrInvalidID privateError = "models: ID is invalid"
	ErrUserIDRequired privateError = "models: user ID is required"
	ErrInvalidPasswordHash privateError = "models: invalid password hash"
	ErrRememberTooShort privateError = "models: invalid remember token"
	ErrRequiredRememberHash privateError = "models: remember has in required"
	ErrUploadFailed privateError = "models: Failed to upload to storage"
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

type privateError string

func (m privateError) Error() string  {
	return string(m)
}