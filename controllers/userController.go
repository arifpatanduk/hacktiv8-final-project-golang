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

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

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

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := utils.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := utils.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, utils.SuccessResponse(gin.H{
		"token": token,
	}, "Login success"))
}