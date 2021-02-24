package models

import (
	"errors"
	"github.com/thiduzz/lenslocked.com/hash"
	"github.com/thiduzz/lenslocked.com/rand"
	"golang.org/x/crypto/bcrypt"
)


type userValidator struct {
	hmac hash.HMAC
	UserRepository
}

type userValidationRule func(*User) error

var _ UserRepository = &userValidator{}

func (uv *userValidator) ByID(id uint) (*User, error) {
	if id <= 0{
		return nil, errors.New("invalid id")
	}
	return uv.UserRepository.ByID(id)
}

func (uv *userValidator) ByRemember(token string) (*User, error) {
	user := User{
		Remember: token,
	}
	if err := validateUserRules(&user, uv.hashRemember); err != nil {
		return nil, err
	}
	return uv.UserRepository.ByRemember(user.RememberHash)
}

func (uv *userValidator) Create(user *User) error {
	if user.Remember == ""{
		token, err := rand.RememberToken()
		if err != nil{
			return err
		}
		user.Remember = token
	}
	if err := validateUserRules(user, uv.bcryptPassword, uv.hashRemember); err != nil {
		return err
	}
	return uv.UserRepository.Create(user)
}

func (uv *userValidator) Update(user *User) error {
	if err := validateUserRules(user, uv.bcryptPassword, uv.hashRemember); err != nil {
		return err
	}
	return uv.UserRepository.Update(user)
}

func (uv *userValidator) Destroy(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	return uv.UserRepository.Destroy(id)
}

func validateUserRules(user *User, fns ... userValidationRule) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

func (uv *userValidator) hashRemember(user *User) error {
	if user.Remember == ""{
		return nil
	}
	user.RememberHash = uv.hmac.Hash(user.Remember)
	return nil
}

func (uv *userValidator) bcryptPassword(user *User) error {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password + userPasswordPepper), bcrypt.DefaultCost)
	if err != nil{
		return ErrInvalidPasswordHash
	}
	user.PasswordHash = string(passwordBytes)
	user.Password = ""
	return nil
}