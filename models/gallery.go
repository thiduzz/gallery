package models

import (
	"gorm.io/gorm"
)

type Gallery struct {
	gorm.Model
	Title  string `gorm:"not null;"`
	Published  bool `gorm:"not null;default:true;"`
	UserID uint
	User User `gorm:"not_null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photos []Photo `gorm:"foreignKey:GalleryID;"`
}

type GalleryService interface {
	GalleryRepository
}

type galleryService struct {
	GalleryRepository
}

func NewGalleryService(db *gorm.DB) GalleryService {
	return &galleryService{
		GalleryRepository: &galleryValidator{
			&galleryGorm{db:db},
		},
	}
}