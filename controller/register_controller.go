package controller

import (
	"backend-golang/database"
	"backend-golang/helpers"
	"backend-golang/models"
	"backend-golang/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	// Initialize request
	var req = structs.UserCreateReq{}

	// Validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Error:   helpers.TranslateErrorMessage(err),
		})
		return
	}

	// create data user
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: helpers.HashPassword(req.Password),
	}

	// save data user
	if err := database.DB.Create(&user).Error; err != nil {
		// Check duplicate entry
		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "User already exists",
				Error:   helpers.TranslateErrorMessage(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create user",
				Error:   helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	// Success response
	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
