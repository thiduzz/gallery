package models

import "gorm.io/gorm"

type PhotoRepository interface {
	Store(photo *Photo) (uint, error)
	Destroy(id uint) error
}


var _ PhotoRepository = &photoGorm{}

type photoGorm struct {
	db *gorm.DB
	BaseRepository
}

func (gg *photoGorm) Destroy(id uint) error {
	return gg.db.Delete(&Photo{}, id).Error
}

func (gg *photoGorm) Store(photo *Photo) (uint, error) {
	result := gg.db.Create(&photo)
	return photo.ID, result.Error
}
