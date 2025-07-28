package controller

import (
	"backend-golang/database"
	"backend-golang/helpers"
	"backend-golang/models"
	"backend-golang/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	// Initialize request
	var req = structs.UserLoginReq{}
	var user = models.User{}

	// Validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Error:   helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Find user by username
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid username",
			Error:   helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Check password

	// If login is success
	token := helpers.GenerateToken(user.Username)

	// Return response
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login success",
		Data: structs.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			Token:     &token,
		},
	})

}
