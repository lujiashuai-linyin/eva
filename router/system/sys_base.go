package system

import (
	v1 "eva/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRouter struct{}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, Response{
				Code: 0,
				Msg:  "success",
			})
		})
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.GET("gt/captcha", baseApi.GetGtCaptcha)
		baseRouter.POST("gt/captcha", baseApi.ValidGtCaptcha)
	}
	return baseRouter
}
