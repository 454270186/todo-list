package handler

import (
	"todolist/dto"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginData dto.LoginReq
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(400, dto.LoginRes{
			StatusCode: -1,
			Msg: "bad request",
		})
		return
	}

	c.JSON(200, dto.LoginRes{
		StatusCode: 0,
		Msg: "login successfully",
	})
}

func Register(c *gin.Context) {

}