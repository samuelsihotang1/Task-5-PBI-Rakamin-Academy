package models

type User struct {
	Model
	Username string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Photos   []Photo
}
