package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        //predefined struct to provide commonly used fields i.e ID ,CreatedAt
	Email      string `gorm:"unique"`
	Password   string
}
