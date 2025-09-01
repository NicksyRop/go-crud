package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model //predefined struct to provide commonly used fields i.e ID ,CreatedAt
	Title      string
	Body       string
}
