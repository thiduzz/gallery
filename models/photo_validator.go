package models

func newPhotoValidator(repo PhotoRepository) *photoValidator {
	return &photoValidator{
		PhotoRepository: repo,
	}
}

type photoValidator struct {
	PhotoRepository
}

type photoValidationRule func(*Photo) error

func validatePhotoRules(photo *Photo, fns ... photoValidationRule) error {
	for _, fn := range fns {
		if err := fn(photo); err != nil {
			return err
		}
	}
	return nil
}

func (gv *photoValidator) Destroy(id uint) error {
	var gallery Gallery
	gallery.ID = id
	if id <= 0 {
		return ErrInvalidID
	}
	return gv.PhotoRepository.Destroy(id)
}

func (gv *photoValidator) titleRequire(photo *Photo) error {
	if photo.Title == ""{
		return ErrRequiredTitle
	}
	return nil
}
