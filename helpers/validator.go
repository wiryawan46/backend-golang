package helpers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"strings"
)

func TranslateErrorMessage(err error) map[string]string {
	// Membuat map untuk menyimpan pesan kesalahan
	errMap := make(map[string]string)

	// Handlign validasi error
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, err := range validationErrors {
			field := err.Field()
			switch err.Tag() {
			case "required":
				errMap[field] = "Field " + field + " is required"
			case "email":
				errMap[field] = "Field " + field + " must be a valid email address"
			case "min":
				errMap[field] = "Field " + field + " must be at least " + err.Param() + " characters long"
			case "max":
				errMap[field] = "Field " + field + " must be at most " + err.Param() + " characters long"
			case "gte":
				errMap[field] = "Field " + field + " must be greater than or equal to " + err.Param()
			case "lte":
				errMap[field] = "Field " + field + " must be less than or equal to " + err.Param()
			default:
				errMap[field] = "Field " + field + " is invalid"
			}
		}
	}

	// Handle gorm error for duplicate entry
	if err != nil {
		// Check if error contain duplicate entry
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errMap["Username"] = "Username already exists"
			}
			if strings.Contains(err.Error(), "email") {
				errMap["Email"] = "Email already exists"
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle gorm error for record not found
			errMap["NotFound"] = "Data not found"
		}
	}

	// Mengembalikan map sebagai JSON
	return errMap
}

// IsDuplicateEntryError detect error from database is duplicate entry
func IsDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}
