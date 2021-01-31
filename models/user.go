package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ErrNotFound = errors.New("models: resource not found")
var ErrInvalidID = errors.New("models: user ID is invalid")
var ErrInvalidPasswordHash = errors.New("models: invalid password hash")
var ErrInvalidPassword = errors.New("models: invalid password")
const userPasswordPepper = "secret-random-string"


type User struct {
	gorm.Model
	Name  string `gorm:"not null;"`
	Email string `gorm:"uniqueIndex:idx_email_unique;not null"`
	Password string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}

func NewUserService(connectionInfo string) (*UserService, error)  {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil{
		return nil, err
	}
	//db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
	return &UserService{
		db: db,
	}, nil
}

type UserService struct{
	db *gorm.DB
}

// Create will create the provided user and backfill data like
// the ID, CreatedAt and UpdatedAt fields
func (u *UserService) Create(user *User) error {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password + userPasswordPepper), bcrypt.DefaultCost)
	if err != nil{
	    return ErrInvalidPasswordHash
	}
	user.PasswordHash = string(passwordBytes)
	user.Password = ""
	return u.db.Create(user).Error
}

// Update will update the provided user will all of the date in
// the provided user object
func (u *UserService) Update(user *User) error {
	return u.db.Save(user).Error
}


// Destroy will softdelete the provided user object from the database
func (u *UserService) Destroy(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	return u.db.Delete(&User{}, id).Error
}

// ByID will look up by the id provided
func (u *UserService) ByID(id uint) (*User, error) {
	var user User
	db := u.db.Where("id = ?", id)
	if err := first(db, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// ByEmail will look up by the email provided
func (u *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := u.db.Where("email = ?", email)
	if err := first(db, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// Authenticate can be used to authenticate a user into the application with a email and password
func (u *UserService) Authenticate(email string, password string) (*User, error) {
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

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound{
		return ErrNotFound
	}
	return err
}