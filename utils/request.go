package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	ErrContentCanNotBeEmpty = errors.New("content can not be empty")
)

func ParseJson(c *gin.Context, payload any) error {
	if c.Request.ContentLength == 0 {
		return ErrContentCanNotBeEmpty
	}

	if err := c.BindJSON(payload); err != nil {
		return err
	}

	return nil
}

func Validate(payload any) map[string]string {
	validate := validator.New()

	err := validate.Struct(payload)

	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMap := make(map[string]string)
			for _, e := range validationErrors {
				errorMap[e.Field()] = fmt.Sprintf("failed on the %s tag", e.Tag())
			}
			return errorMap
		}
	}

	return nil
}

func WriteJson(c *gin.Context, status int, v any) {
	c.JSON(status, map[string](any){
		"data": v,
	})
}

func WriteMessage(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"message": message})
}

func WriteError(c *gin.Context, status int, err error) {
	WriteJson(c, status, gin.H{"error": err.Error()})
}

func WriteServerError(c *gin.Context) {
	WriteJson(c, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
}
