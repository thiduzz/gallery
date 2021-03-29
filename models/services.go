package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Services struct {
	Gallery GalleryService
	User UserService
	Photo PhotoService
	db *gorm.DB
}

func NewServices(connectionInfo string) (*Services, error) {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil{
		return nil, err
	}
	return &Services{
		User:    NewUserService(db),
		Gallery: NewGalleryService(db),
		Photo: NewPhotoService(db),
		db: db,
	}, nil
}

func (sv *Services) AutoMigrate() error {
	return sv.db.AutoMigrate(&User{}, &Gallery{},  &Photo{})
}

func (sv *Services) DestructiveReset() error {
	if err := sv.db.Migrator().DropTable(&User{}, &Gallery{}); err != nil {
		return err
	}
	return sv.AutoMigrate()
}