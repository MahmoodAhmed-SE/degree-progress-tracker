package schemas

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Status     int    `json:"status"`
	Message    string `json:"message"`
	Success    bool   `json:"success"`
	LineNumber int    `json:"line_number,omitempty"`
	FileName   string `json:"file_name,omitempty"`
	Data       any    `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, BaseResponse{
		Status:  statusCode,
		Message: message,
		Success: true,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, err string) {
	_, file, line, _ := runtime.Caller(1)
	c.JSON(statusCode, BaseResponse{
		Status:     statusCode,
		Message:    err,
		Success:    false,
		LineNumber: line,
		FileName:   file,
	})
}
