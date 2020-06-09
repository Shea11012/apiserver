package user

import (
	"apiserver/api/user"
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body user.CreateRequest true "create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"xiaoming"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	zap.L().Info("User create function called.", zap.String("x-request-id", util.GetReqID(c)))

	var r user.CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.User{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := user.CreateRequest{
		Username: r.Username,
	}

	SendResponse(c, nil, rsp)
}
