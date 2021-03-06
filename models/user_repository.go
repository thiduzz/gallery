package models

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	ByID(userID uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	Store(user *User) error
	Update(user *User) error
	Destroy(id uint) error
}

var _ UserRepository = &userGorm{}

type userGorm struct {
	db *gorm.DB
	BaseRepository
}

// Create will create the provided user and backfill data like
// the ID, CreatedAt and UpdatedAt fields
func (ug *userGorm) Store(user *User) error {
	return ug.db.Create(user).Error
}

// Update will update the provided user will all of the date in
// the provided user object
func (ug *userGorm) Update(user *User) error {
	return ug.db.Save(user).Error
}


// Destroy will softdelete the provided user object from the database
func (ug *userGorm) Destroy(id uint) error {
	return ug.db.Delete(&User{}, id).Error
}

// ByID will look up by the id provided
func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id = ?", id)
	if err := ug.first(db, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// ByEmail will look up by the email provided
func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email = ?", email)
	if err := ug.first(db, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// ByRemember will look up by the email provided
func (ug *userGorm) ByRemember(rememberHash string) (*User, error) {
	var user User
	db := ug.db.Where("remember_hash = ?", rememberHash)
	if err := ug.first(db, &user); err != nil {
		return nil, err
	}
	return &user, nil
}