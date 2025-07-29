package controller

import (
	"backend-golang/database"
	"backend-golang/helpers"
	"backend-golang/models"
	"backend-golang/structs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	// Find user by username or email
	if err := database.DB.Where("username = ? OR email = ?", req.Username, req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid username or email",
			Error:   helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Bandingkan password yang dimasukkan dengan password yang sudah di-hash di database
	// Jika tidak cocok, kirimkan respons error Unauthorized
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid Password",
			Error:   helpers.TranslateErrorMessage(err),
		})
		return
	}

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
