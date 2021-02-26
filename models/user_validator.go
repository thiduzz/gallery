package models

import (
	"errors"
	"fmt"
	"github.com/thiduzz/lenslocked.com/hash"
	"github.com/thiduzz/lenslocked.com/rand"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

func newUserValidator(repo UserRepository, hmac hash.HMAC) *userValidator {
	return &userValidator{
		hmac:           hmac,
		UserRepository: repo,
		emailRegex:     regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`),
	}
}

type userValidator struct {
	hmac hash.HMAC
	UserRepository
	emailRegex *regexp.Regexp
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



func (uv *userValidator) ByEmail(email string) (*User, error) {
	user := User{
		Email: email,
	}
	if err := validateUserRules(&user, uv.emailNormalize); err != nil {
		return nil, err
	}
	return uv.UserRepository.ByEmail(user.Email)
}

func (uv *userValidator) Store(user *User) error {
	if err := validateUserRules(user,
		uv.passwordRequire,
		uv.passwordMinLength(8),
		uv.hashPassword,
		uv.passwordHashRequire,

		uv.setDefaultRemember,
		uv.rememberSize(rand.RememberTokenBytes),
		uv.hashRemember,
		uv.hashRememberRequire,

		uv.emailNormalize,
		uv.emailRequire,
		uv.emailFormat,
		uv.emailIsAvailable,
	); err != nil {
		return err
	}
	return uv.UserRepository.Store(user)
}

func (uv *userValidator) Update(user *User) error {
	if err := validateUserRules(user,
		uv.passwordRequire,
		uv.passwordMinLength(8),
		uv.hashPassword,
		uv.passwordHashRequire,

		uv.rememberSize(rand.RememberTokenBytes),
		uv.hashRemember,
		uv.hashRememberRequire,

		uv.emailNormalize,
		uv.emailRequire,
		uv.emailFormat,
		uv.emailIsAvailable,
		); err != nil {
		return err
	}
	return uv.UserRepository.Update(user)
}

func (uv *userValidator) Destroy(id uint) error {
	var user User
	user.ID = id
	if err := validateUserRules(&user, uv.idGreaterThan(0)); err != nil {
		return err
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

func (uv *userValidator) idGreaterThan(n uint) userValidationRule {
	return userValidationRule(func(user *User) error {
		if user.ID <= n{
			return ErrInvalidID
		}
		return nil
	})
}

func (uv *userValidator) emailNormalize(user *User) error {
	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)
	return nil
}

func (uv *userValidator) emailRequire(user *User) error {
	if user.Email == ""{
		return ErrRequiredEmail
	}
	return nil
}

func (uv *userValidator) emailFormat(user *User) error {
	if !uv.emailRegex.MatchString(user.Email){
		return ErrInvalidEmail
	}
	return nil
}

func (uv *userValidator) emailIsAvailable(user *User) error {
	existingUser, err := uv.ByEmail(user.Email)
	if err == ErrNotFound{
		return nil
	}
	if err != nil{
		return err
	}
	if user.ID != existingUser.ID{
		return ErrAlreadyTaken
	}
	return nil
}

func (uv *userValidator) passwordRequire(user *User) error {
	if user.Password == ""{
		return ErrRequiredPassword
	}
	return nil
}

func (uv *userValidator) passwordHashRequire(user *User) error {
	if user.PasswordHash == ""{
		return ErrRequiredPassword
	}
	return nil
}


func (uv *userValidator) passwordMinLength(n uint) userValidationRule {
	return userValidationRule(func(user *User) error {
		if len(user.Password) < int(n){
			return modelError(fmt.Sprintf("Password needs to be at least %d characters",n))
		}
		return nil
	})
}

func (uv *userValidator) hashRemember(user *User) error {
	if user.Remember == ""{
		return nil
	}
	user.RememberHash = uv.hmac.Hash(user.Remember)
	return nil
}

func (uv *userValidator) hashPassword(user *User) error {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password + userPasswordPepper), bcrypt.DefaultCost)
	if err != nil{
		return ErrInvalidPasswordHash
	}
	user.PasswordHash = string(passwordBytes)
	user.Password = ""
	return nil
}


func (uv *userValidator) hashRememberRequire(user *User) error {
	if user.RememberHash == ""{
		return ErrRequiredRememberHash
	}
	return nil
}


func (uv *userValidator) rememberSize(expectedBytes int) userValidationRule {
	return userValidationRule(func(user *User) error {
		bytes, err := rand.NBytes(user.Remember)
		if err != nil{
		    return err
		}
		if bytes < expectedBytes{
			return ErrRememberTooShort
		}
		return nil
	})
}


func (uv *userValidator) setDefaultRemember(user *User) error {
	if user.Remember != ""{
		return nil
	}
	token, err := rand.RememberToken()
	if err != nil{
		return err
	}
	user.Remember = token
	return nil
}