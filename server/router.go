package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/api"
	"github.com/weilaim/blog-api/middleware"
)

//NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	//路由中间件
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	// r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	//路由 
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// // 用户登录
		v1.POST("user/login", api.UserLogin)

		// // 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// // User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}

		//视频操作
		//添加视频
		v1.POST("videos", api.CreateVideo)
		//获取视频
		v1.GET("video/:id",api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		v1.PUT("video/:id", api.UpdateVideo)
		v1.DELETE("video/:id", api.DeleteVideo)

		//排行榜
		v1.GET("rank/daily",api.DailyRank)

		// oss toke
		v1.POST("upload/token",api.UploadToken)
	}
	return r
}