package api

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/conf"
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"

	validator "gopkg.in/go-playground/validator.v8"
)


func Ping(c *gin.Context) {
	c.JSON(200,serializer.Response{
		Status: 0,
		Msg: "网络流畅您可以随意访问········",
	})
}

//获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil{
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}


// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}
