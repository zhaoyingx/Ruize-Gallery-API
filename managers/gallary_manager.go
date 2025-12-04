package managers

import (
	"api-template/models"

	"gorm.io/gorm"
)

type sampleManager struct {
	db *gorm.DB
}
var SampleManager *sampleManager

func (m *sampleManager) QuerySample(university string, year int) ([]models.GraduateEmploymentResponse, error) {
	return nil, nil
}