package user

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 删除一个用户
// @Tag user
// @Accept json
// @Produce json
// @Param id path integer true "用户ID"
// @Security ApiKeyAuth
// @Success 200 {object} api.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [delete]
func Delete(c *gin.Context) {
	userId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := model.DeleteUser(userId); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
