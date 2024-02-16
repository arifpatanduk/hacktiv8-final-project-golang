package models

type SocialMedia struct {
	GormModel
	Name  string `json:"name" form:"name" validate:"required"`
	SocialMediaUrl  string `json:"url" form:"url" validate:"required,url"`
	UserID uint
	User   *User
}
