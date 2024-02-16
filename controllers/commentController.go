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

func GetAllComment(c *gin.Context) {
	// initialize the db
	db := config.GetDB()

	// get all comments from the database
	var comments []models.Comment
	err := db.Find(&comments).Error

	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorResponse("Failed get all comment data"))
		return
	}

	// map the response
	var response []gin.H
	for _, comment := range comments {
		data := gin.H{
			"commentID": comment.ID,
			"userID": comment.UserID,
			"photoID": comment.PhotoID,
			"message": comment.Message,
			"createdAt": comment.CreatedAt,
			"updatedAt": comment.UpdatedAt,
		}
		response = append(response, data)
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully retrieved all comments"))
}

func GetOneComment(c *gin.Context) {
	// initialize the db
	db := config.GetDB()

	commentID := c.Param("commentID")

	// validate that commentID is a valid uint
	parsedCommentID, err := strconv.ParseUint(commentID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid commentID"))
		return
	}

	// get the comment with the given commentID from the database
	var comment models.Comment
	if err := db.First(&comment, parsedCommentID).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Comment not found"))
		return
	}

	// map the response
	response := gin.H{
		"commentID": comment.ID,
		"userID": comment.UserID,
		"photoID": comment.PhotoID,
		"message": comment.Message,
		"createdAt": comment.CreatedAt,
		"updatedAt": comment.UpdatedAt,
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully retrieved comment data"))
}

func CreateComment(c *gin.Context) {
	// initialize the db
	db := config.GetDB()

	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	Comment := models.Comment{}
	
	// get logged userID
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// validate json
	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	// validate input 
	err := utils.ValidateStruct(Comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.TranslateValidationErrors(err)))
		return
	}

	// store comment to db
	Comment.UserID = userID
	errCreate := db.Debug().Create(&Comment).Error
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create comment data"))
		return
	}

	// map the response
	response := gin.H{
		"commentID": Comment.ID,
		"userID": Comment.UserID,
		"photoID": Comment.PhotoID,
		"message": Comment.Message,
		"createdAt": Comment.CreatedAt,
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully created the comment data"))
}

func UpdateComment(c *gin.Context) {
	// initialize the db
	db := config.GetDB()

	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	CommentInput := models.CommentInput{}

	// get commentID
	commentID := c.Param("commentID")

	// validate json
	if contentType == appJSON {
		c.ShouldBindJSON(&CommentInput)
	} else {
		c.ShouldBind(&CommentInput)
	}

	// validate input 
	err := utils.ValidateStruct(CommentInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.TranslateValidationErrors(err)))
		return
	}

	// validate that commentID is a valid uint
	parsedCommentID, err := strconv.ParseUint(commentID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid commentID"))
		return
	}

	// check if the comment is exists
	var comment models.Comment
	if err := db.First(&comment, parsedCommentID).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Comment not found"))
		return
	}

	// update the comment data
	comment.Message = CommentInput.Message
	db.Save(&comment)

	// map the response
	response := gin.H{
		"commentID": comment.ID,
		"userID": comment.UserID,
		"photoID": comment.PhotoID,
		"message": comment.Message,
		"updatedAt": comment.UpdatedAt,
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Successfully updated the comment data"))
}

func DeleteComment(c *gin.Context) {
	// initialize the db
	db := config.GetDB()
	commentID := c.Param("commentID")

	// validate that commentID is a valid uint
	parsedCommentID, err := strconv.ParseUint(commentID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid commentID"))
		return
	}

	// cek if the comment is exists
	var comment models.Comment
	if err := db.First(&comment, parsedCommentID).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Comment not found"))
		return
	}

	// delete the comment from the database
	if err := db.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete comment"))
		return
	}

	// return response
	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Successfully deleted comment data for commentID: " + commentID))
}