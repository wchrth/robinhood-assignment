package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, apiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func RespondSuccessNoData(c *gin.Context, code int, message string) {
	c.JSON(code, apiResponse{
		Status:  "success",
		Message: message,
	})
}

func RespondError(c *gin.Context, err error) {
	if e, ok := err.(*APIError); ok {
		c.JSON(e.Code, apiResponse{
			Status:  "error",
			Message: e.Message,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, apiResponse{
		Status:  "error",
		Message: http.StatusText(http.StatusInternalServerError),
	})
}
