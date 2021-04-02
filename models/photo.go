package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title  string `gorm:"not null;"`
	Filename string `gorm:"default:null;not null;"`
	Path  string `gorm:"not null;"`
	GalleryID uint `gorm:"not null;"`
}

type PhotoService interface {
	PhotoRepository
}

type photoService struct {
	PhotoRepository
}

func NewPhotoService(db *gorm.DB) PhotoService {
	return &photoService{
		PhotoRepository: &photoValidator{
			&photoGorm{db:db},
		},
	}
}