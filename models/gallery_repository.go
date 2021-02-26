package models

import "gorm.io/gorm"

type GalleryRepository interface {
	Store(gallery *Gallery) error
}


var _ GalleryRepository = &galleryGorm{}

type galleryGorm struct {
	db *gorm.DB
	BaseRepository
}

func (gg *galleryGorm) Store(gallery *Gallery) error {
	return gg.db.Create(gallery).Error
}

func (gg *galleryGorm) Update(gallery *Gallery) error {
	return gg.db.Save(gallery).Error
}