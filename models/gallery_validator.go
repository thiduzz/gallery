package models

import (
	"fmt"
)

func newGalleryValidator(repo GalleryRepository) *galleryValidator {
	return &galleryValidator{
		GalleryRepository: repo,
	}
}

type galleryValidator struct {
	GalleryRepository
}

type galleryValidationRule func(*Gallery) error

func validateGalleryRules(gallery *Gallery, fns ... galleryValidationRule) error {
	for _, fn := range fns {
		if err := fn(gallery); err != nil {
			return err
		}
	}
	return nil
}

func (gv *galleryValidator) Store(gallery *Gallery) error {
	if err := validateGalleryRules(gallery,
		gv.titleRequire,
		gv.titleMinLength(3),
		gv.userIDRequired,
	); err != nil {
		return err
	}
	return gv.GalleryRepository.Store(gallery)
}


func (gv *galleryValidator) Update(gallery *Gallery) error {
	if err := validateGalleryRules(gallery,
		gv.titleRequire,
		gv.titleMinLength(3),
		gv.userIDRequired,
	); err != nil {
		return err
	}
	return gv.GalleryRepository.Update(gallery)
}

func (gv *galleryValidator) Destroy(id uint) error {
	var gallery Gallery
	gallery.ID = id
	if id <= 0 {
		return ErrInvalidID
	}
	return gv.GalleryRepository.Destroy(id)
}

func (gv *galleryValidator) titleRequire(gallery *Gallery) error {
	if gallery.Title == ""{
		return ErrRequiredTitle
	}
	return nil
}

func (gv *galleryValidator) userIDRequired(gallery *Gallery) error {
	if gallery.UserID <= 0{
		return ErrUserIDRequired
	}
	return nil
}

func (gv *galleryValidator) titleMinLength(n uint) galleryValidationRule {
	return galleryValidationRule(func(gallery *Gallery) error {
		if gallery.Title == "" {
			return nil
		}

		if len(gallery.Title) < int(n){
			return modelError(fmt.Sprintf("Title needs to be at least %d characters",n))
		}
		return nil
	})
}
