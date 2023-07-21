package main

import (
	"todolist/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	api := r.Group("/api")

	api.POST("/login", handler.Login)
	api.POST("/register", handler.Register)

	return r
}