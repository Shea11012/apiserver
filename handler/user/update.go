package user

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// @Summary 更新用户信息
// @Security ApiKeyAuth
// @Tag user
// @Accept json
// @Produce json
// @Param id path integer true "更新用户 id"
// @Param user body model.User true "需要更新的用户信息"
// @Success 200 {object} api.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [put]
func Update(c *gin.Context) {
	zap.L().Info("update function called", zap.String("x-Request-Id", util.GetReqID(c)))

	userId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var u model.User
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.Id = userId
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
