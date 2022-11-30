package system

import (
	"eva/biz/api/v1"
	"eva/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithJWT := Router.Group("user").Use(middleware.JWTAuth())
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	//{
	//userRouter.POST("adminRegister", baseApi.Register)         // 管理员注册账号
	userRouterWithoutRecord.POST("/register", baseApi.Register)       // 注册账号
	userRouterWithoutRecord.GET("/mobile/:phone", baseApi.CheckPhone) // 验证手机号
	userRouterWithoutRecord.GET("/sms/valid", baseApi.SendMessage)    // 发送短信
	//userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
	//userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
	//userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
	//userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
	//userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
	//	userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
	//	userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
	//}
	//{
	//	userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
	userRouterWithJWT.GET("/getUserInfo", baseApi.GetUserInfo) // 获取自身信息
	//}
}
