package zgin

import "github.com/gin-gonic/gin"

type IProtect interface {
	InRequest(*gin.Context) error
}
