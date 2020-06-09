package user

import (
	"apiserver/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func Relative(c *gin.Context) {
	userid,_ := strconv.ParseUint(c.Param("userid"),10,64)
	card := model.Card{
		UserId:userid,
		Name:   c.Param("name"),
	}
	if err := card.Create(); err != nil {
		zap.S().Error(err)
	}
	c.JSON(200,card)
}

func GetRelative(c *gin.Context) {
	userid,_ := strconv.ParseUint(c.Param("id"),10,64)
	u,err := model.GetUserById(userid)
	if err != nil {
		c.JSON(200,err)
		return
	}

	cards,err := u.GetCards()
	if err != nil {
		c.JSON(200,err)
		return
	}
	c.JSON(200,cards)
}
