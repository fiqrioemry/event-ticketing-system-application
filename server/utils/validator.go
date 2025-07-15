package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/fiqrioemry/go-api-toolkit/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidateJSON[T any](c *gin.Context, req *T) bool {
	if err := c.ShouldBindJSON(req); err != nil {

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			validationErr := buildValidationError(validationErrors)
			response.Error(c, validationErr)
			return false
		}

		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			parseErr := response.NewBadRequest("Invalid data type for field").
				WithContext("field", jsonErr.Field).
				WithContext("expected_type", jsonErr.Type.String())
			response.Error(c, parseErr)
			return false
		}

		if syntaxErr, ok := err.(*json.SyntaxError); ok {
			parseErr := response.NewBadRequest("Invalid JSON syntax").
				WithContext("offset", syntaxErr.Offset)
			response.Error(c, parseErr)
			return false
		}

		parseErr := response.NewBadRequest("Invalid JSON format")
		response.Error(c, parseErr)
		return false
	}
	return true
}

// Enhanced BindAndValidateForm with better error logging
func BindAndValidateForm[T any](c *gin.Context, req *T) bool {
	if err := c.ShouldBind(req); err != nil {
		// Debug: Log the raw error
		log.Printf("Bind error: %v", err)
		log.Printf("Request content type: %s", c.GetHeader("Content-Type"))

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			log.Printf("Validation errors: %+v", validationErrors)
			validationErr := buildValidationError(validationErrors)
			response.Error(c, validationErr)
			return false
		}

		// Log the specific binding error
		log.Printf("Form binding failed: %v", err)
		// ✅ FIXED: Use response package instead of errors package
		formErr := response.NewBadRequest(fmt.Sprintf("Invalid form data format: %v", err))
		response.Error(c, formErr)
		return false
	}
	return true
}

// ✅ FIXED: Use response.AppError instead of errors.AppError
func buildValidationError(validationErrors validator.ValidationErrors) *response.AppError {
	errorDetails := make(map[string]any)

	for _, fieldError := range validationErrors {
		fieldName := strings.ToLower(fieldError.Field())

		switch fieldError.Tag() {
		case "required":
			errorDetails[fieldName] = fmt.Sprintf("%s is required", fieldName)
		case "email":
			errorDetails[fieldName] = "Please provide a valid email address"
		case "min":
			errorDetails[fieldName] = fmt.Sprintf("%s must be at least %s characters", fieldName, fieldError.Param())
		case "max":
			errorDetails[fieldName] = fmt.Sprintf("%s must be at most %s characters", fieldName, fieldError.Param())
		case "len":
			errorDetails[fieldName] = fmt.Sprintf("%s must be exactly %s characters", fieldName, fieldError.Param())
		case "numeric":
			errorDetails[fieldName] = fmt.Sprintf("%s must be numeric", fieldName)
		case "alpha":
			errorDetails[fieldName] = fmt.Sprintf("%s must contain only letters", fieldName)
		case "alphanum":
			errorDetails[fieldName] = fmt.Sprintf("%s must contain only letters and numbers", fieldName)
		case "url":
			errorDetails[fieldName] = fmt.Sprintf("%s must be a valid URL", fieldName)
		case "uuid":
			errorDetails[fieldName] = fmt.Sprintf("%s must be a valid UUID", fieldName)
		default:
			errorDetails[fieldName] = fmt.Sprintf("%s is invalid", fieldName)
		}
	}

	// ✅ FIXED: Use response package and proper return
	return response.NewBadRequest("Validation failed").WithContext("errors", errorDetails)
}

func ValidateStruct(s any) error {
	validate := validator.New()
	return validate.Struct(s)
}
