package models

import "gorm.io/gorm"

type BaseRepository struct {}

func (receiver *BaseRepository) first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound{
		return ErrNotFound
	}
	return err
}
