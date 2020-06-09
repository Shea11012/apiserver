package user

import (
	"apiserver/api/user"
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service"
	"github.com/gin-gonic/gin"
)

// @Summary 用户列表
// @Tag user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} user.ListResponse "{"code":0,"message":"OK","data":{"totalCount":10,"userList":[{"id":0,"username":"xiaoming","sayHello":"Hello,xxx","password":"xxx","createdAt":"2020-06-09 12:31:21","updateAt":"2020-06-09 12:31:12"}]}}"
// @Router /user [get]
func List(c *gin.Context) {
	var r user.ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)

	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, user.ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
