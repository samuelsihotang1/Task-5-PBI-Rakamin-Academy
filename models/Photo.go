package models

import (
	"gorm.io/gorm"
)

type Photo struct {
    gorm.Model
    Title    string 
    Caption  string 
    PhotoUrl string 
    UserID   uint   
    User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
