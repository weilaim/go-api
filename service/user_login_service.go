package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

//UserLoginService 管理用户注册服务
type UserLoginService struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
}


//设置session
func (server *UserLoginService)setSeesion(c *gin.Context, user model.User){
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()

}
//用户登录 Login
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name = ?",service.UserName).First(&user).Error; err != nil{
		return serializer.ParamErr("账号或密码错误",nil)
	}
	
	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码",nil)
	}

	//设置session 
	service.setSeesion(c, user)
	
	return serializer.BuildUserResponse(user)
}

