package models

import "gorm.io/gorm"

type GraduateEmployment struct {
	gorm.Model
	Year                   int     `json:"year" gorm:"index"`
	University             string  `json:"university" gorm:"index;size:255"`
	School                 string  `json:"school" gorm:"size:255"`
	Degree                 string  `json:"degree" gorm:"size:255"`
	EmploymentRateOverall  float64 `json:"employment_rate_overall"`
	EmploymentRateFtPerm   float64 `json:"employment_rate_ft_perm"`
	BasicMonthlyMean       float64 `json:"basic_monthly_mean"`
	BasicMonthlyMedian     float64 `json:"basic_monthly_median"`
	GrossMonthlyMean       float64 `json:"gross_monthly_mean"`
	GrossMonthlyMedian     float64 `json:"gross_monthly_median"`
	GrossMthly25Percentile float64 `json:"gross_mthly_25_percentile"`
	GrossMthly75Percentile float64 `json:"gross_mthly_75_percentile"`
}

func (GraduateEmployment) TableName() string {
	return "graduate_employment"
}

type GraduateEmploymentResponse struct {
	School                string `json:"school"`
	EmploymentRateOverall string `json:"employment_rate_overall"`
}

type Images struct {
	Key string `json:"key"`
    Url string `json:"url"`
}
