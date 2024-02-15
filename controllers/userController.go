package controllers

import (
	"go-mygram/config"
	"go-mygram/models"
	"go-mygram/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func RegisterUser(c *gin.Context) {
	db := config.GetDB()
	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	// validate json
	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	// validate input 
	err := utils.ValidateStruct(User)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.TranslateValidationErrors(err)))
		return
	}

	// store to db
	errCreate := db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(errCreate.Error()))
		return
	}

	// success response
	response := gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"username": User.Username,
		"age": User.Age,
	}
	c.JSON(http.StatusCreated, utils.SuccessResponse(response, "Register success"))
}

func LoginUser(c *gin.Context) {
	db := config.GetDB()
	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	// validate json
	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	// check the user according to the email
	password = User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	// check the password
	comparePass := utils.ComparePass([]byte(User.Password), []byte(password))

	// error response
	if (err != nil || !comparePass) {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("invalid email or password"))
		return
	}

	// generate token
	token := utils.GenerateToken(User.ID, User.Email)

	// success response
	c.JSON(http.StatusOK, utils.SuccessResponse(gin.H{
		"token": token,
	}, "Login success"))
}