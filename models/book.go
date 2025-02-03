package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" validate:"required,min=2"`  // add tags for validation
	Author string `json:"author" validate:"required,min=3"` // add tags for validation
}
