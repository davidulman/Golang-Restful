package models

import "gorm.io/gorm"

type User struct {
    gorm.Model

    FirstName string `gorm:"not null" json:"firstName"`
    LastName  string `gorm:"not null" json:"lastName"`
    Email     string `gorm:"unique" json:"email"`
    Password  string `gorm:"not null" json:"password"`
}
