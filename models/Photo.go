package models

type Photo struct {
	Model
	Title    string
	Caption  string
	PhotoUrl string
	UserID   uint
	User     User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
