package middleware

import (
	"eva/global"
	"eva/model/common/response"
	"eva/service"
	"eva/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.AuthorityId))
		e := casbinService.Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.EVA_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
