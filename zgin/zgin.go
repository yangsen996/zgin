package zgin

import "github.com/gin-gonic/gin"

type ZGin struct {
	*gin.Engine
	group *gin.RouterGroup
}

func NewZGin() *ZGin {
	return &ZGin{Engine: gin.New()}
}

// Launch 启动函数
func (zgin *ZGin) Launch() {
	zgin.Run(":9999")
}

// Formation controllers统一入口
func (zgin *ZGin) Formation(groupName string, builds ...IBuildInterface) *ZGin {
	// 添加组
	zgin.group = zgin.Group(groupName)
	//同一的api接口
	for _, build := range builds {
		build.Build(zgin)
	}
	return zgin
}

// Handle 重写handle添加group
func (zgin *ZGin) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *ZGin {
	zgin.group.Handle(httpMethod, relativePath, handlers...)
	return zgin
}

//Middlewares 中间件
func (zgin *ZGin) Middlewares(m IProtect) *ZGin {
	zgin.Use(func(c *gin.Context) {
		err := m.InRequest(c)
		if err != nil {
			c.Abort()
		} else {
			c.Next()
		}
	})
	return zgin
}
