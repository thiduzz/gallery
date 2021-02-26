package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	GalleryID uint
	Name  string `gorm:"not null;"`
	Path  string `gorm:"not null;"`
}