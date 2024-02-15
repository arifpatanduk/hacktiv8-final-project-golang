package models

import (
	"go-mygram/utils"

	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" validate:"required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required,email"`
	Password     string        `gorm:"not null" json:"password" form:"password" validate:"required,min=6"`
	Age          int           `gorm:"not null" json:"age" form:"age" validate:"min=8"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = utils.HashPass(u.Password)
	err = nil
	return
}
