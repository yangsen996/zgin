package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
)

type UserMiddleware struct{}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (u *UserMiddleware) InRequest(c *gin.Context) error {
	log.Println("用户中间件")
	return nil
}
