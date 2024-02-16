package models

type Comment struct {
	GormModel
	Message string `json:"message" form:"message" validate:"required"`
	PhotoID uint   `json:"photoID" form:"photoID" validate:"required,number"`
	Photo   *Photo
	UserID  uint
	User    *User
}

type CommentInput struct {
	Message string `json:"message" form:"message" validate:"required"`
}
