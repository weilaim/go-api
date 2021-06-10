package api

import (
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/service"
)

//CreateVideo 用户注册接口
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}