package controller

import (
	"backend-golang/database"
	"backend-golang/helpers"
	"backend-golang/models"
	"backend-golang/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindUser(c *gin.Context) {
	//initialize model
	var users []models.User

	// find user data
	database.DB.Find(&users)

	// return user data
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Lists of users",
		Data:    users,
	})
}

func CreateUser(c *gin.Context) {
	{

		//initialize request
		var req = structs.UserCreateReq{}

		//validate request
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
				return
			}
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create user",
				Error:   helpers.TranslateErrorMessage(err),
			})
			return
		}
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
}

func FindByUserId(c *gin.Context) {
	//get param id
	id := c.Param("id")

	//initialize model
	var users models.User

	// find user data BY ID
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Error:   helpers.TranslateErrorMessage(err),
		})
		return
	}

	// return user data
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User found",
		Data: structs.UserResponse{
			ID:        users.ID,
			Name:      users.Name,
			Username:  users.Username,
			Email:     users.Email,
			Password:  users.Password,
			CreatedAt: users.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: users.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateUser(c *gin.Context) {
	{

		// get id
		id := c.Param("id")

		//find user first
		var user models.User
		if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, structs.ErrorResponse{
				Success: false,
				Message: "User not found",
				Error:   helpers.TranslateErrorMessage(err),
			})
			return
		}

		//initialize request
		var req = structs.UserUpdateReq{}

		//validate request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
				Success: false,
				Message: "Validation error",
				Error:   helpers.TranslateErrorMessage(err),
			})
			return
		}
		// create data user
		user.Name = req.Name
		user.Username = req.Username
		user.Email = req.Email
		user.Password = helpers.HashPassword(req.Password)

		// save data user
		if err := database.DB.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to update user",
				Error:   helpers.TranslateErrorMessage(err),
			})
			return
		}
		c.JSON(http.StatusCreated, structs.SuccessResponse{
			Success: true,
			Message: "User updated successfully",
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
}

func DeleteUser(c *gin.Context) {
	// get id
	id := c.Param("id")

	//find user first
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Error:   helpers.TranslateErrorMessage(err),
		})
		return
	}

	// delete user
	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to delete user",
			Error:   helpers.TranslateErrorMessage(err),
		})
		return
	}
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}
