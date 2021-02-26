package models

import (
	"github.com/thiduzz/lenslocked.com/hash"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const userPasswordPepper = "secret-random-string"
const hmacSecretKey = "secret-hmac-key"

type User struct {
	gorm.Model
	Name  string `gorm:"not null;"`
	Email string `gorm:"uniqueIndex:idx_email_unique;not null"`
	Password string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember string `gorm:"-"`
	RememberHash string `gorm:"not_null;uniqueIndex"`
}

type UserService interface {
	Authenticate(email string, password string) (*User, error)
	UserRepository
}

type userService struct {
	UserRepository
}

var _ UserService = &userService{}

func NewUserService(db *gorm.DB) UserService  {
	ug :=  &userGorm{db: db}
	return &userService{
		UserRepository: newUserValidator(ug, hash.NewHMAC(hmacSecretKey)),
	}
}


// Authenticate can be used to authenticate a user into the application with a email and password
func (u *userService) Authenticate(email string, password string) (*User, error) {
	foundUser, err := u.ByEmail(email)
	if err != nil{
	    return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password + userPasswordPepper))
	switch err {
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrInvalidPassword
	case nil:
	default:
		return nil, err
	}
	return foundUser, nil
}