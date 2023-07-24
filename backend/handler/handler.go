package handler

import (
	"log"
	"strconv"
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

	userID, err := serviceHandler.Login(loginData.Username, loginData.Password)
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
		UserID: userID,
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

	newUserID, err := serviceHandler.Register(registerData.Username, registerData.Password)
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
		UserID: newUserID,
	})
}

func GetTasksByID(c *gin.Context) {
	userIDstr := c.Param("user_id")
	if userIDstr == "" {
		c.JSON(400, gin.H {
			"status_code": -1,
			"msg": "user id cannot be empty",
		})
		return
	}

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.JSON(400, gin.H{
			"status_code": -1,
			"msg": "user id must be integer",
		})
		return
	}

	tasks, err := serviceHandler.GetTasksByID(userID)
	if err != nil {
		c.JSON(500, gin.H{
			"status_code": -1,
			"msg": "internal error: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": 0,
		"msg": "successful",
		"tasks": tasks,
	})
}

func AddTask(c *gin.Context) {
	taskID := c.Query("task_id")
	taskName := c.Query("name")
	userIDstr := c.Query("user_id")
	if userIDstr == "" {
		c.JSON(400, gin.H {
			"status_code": -1,
			"msg": "user id cannot be empty",
		})
		return
	}

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.JSON(400, gin.H{
			"status_code": -1,
			"msg": "user id must be integer",
		})
		return
	}

	err = serviceHandler.AddNewTask(userID, taskName, taskID)
	if err != nil {
		c.JSON(400, gin.H {
			"status_code": -1,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H {
		"status_code": 0,
		"msg": "add task successfully",
	})
}

func DelTask(c *gin.Context) {
	taskId := c.Query("task_id")
	if taskId == "" {
		c.JSON(400, gin.H {
			"status_code": -1,
			"msg": "task id cannot be empty",
		})
		return
	}

	err := serviceHandler.DelTaskByID(taskId)
	if err != nil {
		c.JSON(400, gin.H {
			"status_code": -1,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H {
		"status_code": 0,
		"msg": "delete task successfully",
	})
}