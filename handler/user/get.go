package user

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
)

// @Summary 获取用户的详细信息
// @Tag user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param username path string true "用户名"
// @Success 200 {object} model.User "{"code":0,"message":"OK","data":{"username":"xiaoming","password":"xxxx"}}"
// @Router /user/{username} [get]
func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, user)
}
