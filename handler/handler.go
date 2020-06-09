package handler

import (
	"apiserver/api"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, api.Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
