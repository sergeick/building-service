package models

type Building struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	City        string `json:"city" validate:"required"`
	YearBuilt   int    `json:"year" validate:"required"`
	FloorsCount int    `json:"floors" validate:"required"`
}
