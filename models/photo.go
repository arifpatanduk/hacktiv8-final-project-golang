package models

import "mime/multipart"

type Photo struct {
	GormModel
	Title    string `json:"title" form:"title" validdate:"required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" validdate:"required"`
	UserID   uint
	User     *User
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

type PhotoInput struct {
	Title   string           `json:"title" form:"title" validate:"required"`
	Caption string           `json:"caption" form:"caption"`
	Photo   *multipart.FileHeader `json:"photo" form:"photo" validate:"required, image"`
}
