package services

import (
	"api-template/models"
)

var GalleryService = &galleryService{}

type galleryService struct{}

func (s *galleryService) QueryByYear(year int) ([]models.Images, error) {
	return nil, nil
}
