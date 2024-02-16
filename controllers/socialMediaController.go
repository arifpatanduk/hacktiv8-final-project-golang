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

func GetAllSocialMedia(c *gin.Context) {
	// initialize the db
	db := config.GetDB()

	// get all socialMedia from the database
	var socialMedia []models.SocialMedia
	err := db.Find(&socialMedia).Error

	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorResponse("Failed get all social media"))
		return
	}

	// map the response
	var response []gin.H
	for _, sosmed := range socialMedia {
		data := gin.H{
			"sosmedID": sosmed.ID,
			"userID": sosmed.UserID,
			"name": sosmed.Name,
			"url": sosmed.SocialMediaUrl,
			"createdAt": sosmed.CreatedAt,
			"updatedAt": sosmed.UpdatedAt,
		}
		response = append(response, data)
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully retrieved all Social Media"))
}

func GetOneSocialMedia(c *gin.Context) {
	// initialize the db
	db := config.GetDB()

	sosmedID := c.Param("sosmedID")

	// validate that sosmedID is a valid uint
	parsedSosmedID, err := strconv.ParseUint(sosmedID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid sosmedID"))
		return
	}

	// get the social media with the given sosmedID from the database
	var sosmed models.SocialMedia
	if err := db.First(&sosmed, parsedSosmedID).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Social Media not found"))
		return
	}

	// map the response
	response := gin.H{
		"sosmedID": sosmed.ID,
		"userID": sosmed.UserID,
		"name": sosmed.Name,
		"url": sosmed.SocialMediaUrl,
		"createdAt": sosmed.CreatedAt,
		"updatedAt": sosmed.UpdatedAt,
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully retrieved social media data"))
}

func CreateSocialMedia(c *gin.Context) {
	// initialize the db
	db := config.GetDB()

	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}
	
	// get logged userID
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// validate json
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	// validate input 
	err := utils.ValidateStruct(SocialMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.TranslateValidationErrors(err)))
		return
	}

	// store social media to db
	SocialMedia.UserID = userID
	errCreate := db.Debug().Create(&SocialMedia).Error
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create social media data"))
		return
	}

	// map the response
	response := gin.H{
		"sosmedID": SocialMedia.ID,
		"userID": SocialMedia.UserID,
		"name": SocialMedia.Name,
		"url": SocialMedia.SocialMediaUrl,
		"createdAt": SocialMedia.CreatedAt,
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully created the social media data"))
}

func UpdateSocialMedia(c *gin.Context) {
	// initialize the db
	db := config.GetDB()

	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	SocialMediaInput := models.SocialMedia{}

	// get sosmedID
	sosmedID := c.Param("sosmedID")

	// validate json
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMediaInput)
	} else {
		c.ShouldBind(&SocialMediaInput)
	}

	// validate input 
	err := utils.ValidateStruct(SocialMediaInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.TranslateValidationErrors(err)))
		return
	}

	// validate that sosmedID is a valid uint
	parsedSosmedID, err := strconv.ParseUint(sosmedID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid sosmedID"))
		return
	}

	// check if the social media is exists
	var existingSosmed models.SocialMedia
	if err := db.First(&existingSosmed, parsedSosmedID).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Social Media not found"))
		return
	}

	// update the social media data
	existingSosmed.Name = SocialMediaInput.Name
	existingSosmed.SocialMediaUrl = SocialMediaInput.SocialMediaUrl
	db.Save(&existingSosmed)

	// map the response
	response := gin.H{
		"sosmedID": existingSosmed.ID,
		"userID": existingSosmed.UserID,
		"name": existingSosmed.Name,
		"url": existingSosmed.SocialMediaUrl,
		"updatedAt": existingSosmed.UpdatedAt,
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully updated the social media data"))
}

func DeleteSocialMedia(c *gin.Context) {
	// initialize the db
	db := config.GetDB()
	sosmedID := c.Param("sosmedID")

	// validate that sosmedID is a valid uint
	parsedSosmedID, err := strconv.ParseUint(sosmedID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid sosmedID"))
		return
	}

	// cek if the social media is exists
	var existingSosmed models.SocialMedia
	if err := db.First(&existingSosmed, parsedSosmedID).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Photo not found"))
		return
	}

	// delete the photo from the database
	if err := db.Delete(&existingSosmed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete social media"))
		return
	}

	// return response
	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Successfully deleted social media data for sosmedID: " + sosmedID))
}