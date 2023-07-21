package handler

import (
	"log"
	"todolist/dto"
	"todolist/service"

	"github.com/gin-gonic/gin"
)

var serviceHandler *service.ServiceHandler

func init() {
	serviceHandler = service.NewServiceHandler()
}

func Login(c *gin.Context) {
	var loginData dto.LoginReq
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(400, dto.LoginRes{
			StatusCode: -1,
			Msg: "bad request",
		})
		return
	}

	err := serviceHandler.Login(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(400, dto.LoginRes{
			StatusCode: -1,
			Msg: err.Error(),
		})
		return
	}

	c.JSON(200, dto.LoginRes{
		StatusCode: 0,
		Msg: "login successfully",
	})
}

func Register(c *gin.Context) {
	var registerData dto.RegisterReq
	if err := c.BindJSON(&registerData); err != nil {
		c.JSON(400, dto.RegisterRes{
			StatusCode: -1,
			Msg: "bad request",
		})
		return
	}

	err := serviceHandler.Register(registerData.Username, registerData.Password)
	if err != nil {
		c.JSON(400, dto.RegisterRes{
			StatusCode: -1,
			Msg: err.Error(),
		})
		log.Println(err)
		return
	}

	c.JSON(200, dto.RegisterRes{
		StatusCode: 0,
		Msg: "register successfully",
	})
}