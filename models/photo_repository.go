package models

import "gorm.io/gorm"

type PhotoRepository interface {
	Store(photo *Photo) error
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

func (gg *photoGorm) Store(photo *Photo) error {
	return gg.db.Create(photo).Error
}
