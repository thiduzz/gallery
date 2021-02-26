package models

func newGalleryValidator(repo GalleryRepository) *galleryValidator {
	return &galleryValidator{
		GalleryRepository: repo,
	}
}

type galleryValidator struct {
	GalleryRepository
}
