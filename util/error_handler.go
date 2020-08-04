package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var errorMap = map[string]int{
	"ID is invalid":       http.StatusBadRequest,
	"Name is invalid":     http.StatusBadRequest,
	"Username is invalid": http.StatusBadRequest,
	"Password is invalid": http.StatusBadRequest,
}

// HandleError handle error in response
func HandleError(c *gin.Context, errCode int, err error) {
	statusCode := errorMap[err.Error()]

	c.JSON(statusCode, gin.H{
		"code":    errCode,
		"message": err.Error(),
	})
}
