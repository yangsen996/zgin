package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yangsen996/zgin/zgin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (user *UserController) UserList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"result": "user list success"})
	}
}

func (user *UserController) Build(zgin *zgin.ZGin) {
	zgin.Handle("GET", "user_list", user.UserList())
}
