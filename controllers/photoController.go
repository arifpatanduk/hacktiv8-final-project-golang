package controllers

import (
	"go-mygram/config"
	"go-mygram/models"
	"go-mygram/utils"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAllPhoto(c *gin.Context) {
	// initialize the db
	db := config.GetDB()

	// get all photos from the database
	var photos []models.Photo
	err := db.Find(&photos)

	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorResponse("Failed get all photos"))
		return
	}

	// map the response
	var response []gin.H
	for _, photo := range photos {
		data := gin.H{
			"photoID": photo.ID,
			"title": photo.Title,
			"caption": photo.Caption,
			"photoUrl": photo.PhotoUrl,
			"createdAt": photo.CreatedAt,
			"updatedAt": photo.UpdatedAt,
		}
		response = append(response, data)
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully retrieved all photos"))
}

func GetOnePhoto(c *gin.Context) {
	// initialize the db
	db := config.GetDB()

	photoID := c.Param("photoID")

	// validate that photoID is a valid uint
	parsedphotoID, err := strconv.ParseUint(photoID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid photoID"))
		return
	}

	// get the photo with the given photoID from the database
	var photo models.Photo
	if err := db.First(&photo, parsedphotoID).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Photo not found"))
		return
	}

	// map the response
	response := gin.H{
		"photoID": photo.ID,
		"title": photo.Title,
		"caption": photo.Caption,
		"photoUrl": photo.PhotoUrl,
		"createdAt": photo.CreatedAt,
		"updatedAt": photo.UpdatedAt,
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully retrieved photo data"))
}

func CreatePhoto(c *gin.Context) {

	// initialize the db
	db := config.GetDB()

	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	PhotoInput := models.PhotoInput{}
	
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// validate json
	if contentType == appJSON {
		c.ShouldBindJSON(&PhotoInput)
	} else {
		c.ShouldBind(&PhotoInput)
	}

	// upload photo to cloudinary and get the url
	cloudinaryURL, errUpload := utils.UploadToCloudinary(PhotoInput.Photo, "photos")
	if errUpload != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to upload photo"))
		return
	}

	// create photo object
	photo := models.Photo{
		Title:    PhotoInput.Title,
		Caption:  PhotoInput.Caption,
		PhotoUrl: cloudinaryURL,
		UserID:   userID,
	}

	// store photo to db
	errCreate := db.Debug().Create(&photo).Error
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create photo data"))
		return
	}

	// map the response
	response := gin.H{
		"photoID": photo.ID,
		"title": photo.Title,
		"caption": photo.Caption,
		"photoUrl": photo.PhotoUrl,
		"createdAt": photo.CreatedAt,
		"updatedAt": photo.UpdatedAt,
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully created the photo data"))
}

func UpdatePhoto(c *gin.Context) {
	// 
}

func DeletePhoto(c *gin.Context) {
	// 
}