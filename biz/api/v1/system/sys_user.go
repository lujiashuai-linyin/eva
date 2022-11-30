package system

import (
	"context"
	utils2 "eva/biz/utils"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"eva/global"
	"eva/model/common/request"
	"eva/model/common/response"
	"eva/model/system"
	systemReq "eva/model/system/request"
	systemRes "eva/model/system/response"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils2.Verify(l, utils2.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &system.SysUser{Username: l.Username, Password: l.Password}
	user, err := userService.Login(u)
	if err != nil {
		response.FailWithMessage("登录失败", c)
	}
	b.TokenNext(c, *user)
	return
	//if store.Verify(l.CaptchaId, l.Captcha, true) {
	//	u := &system.SysUser{Username: l.Username, Password: l.Password}
	//	user, err := userService.Login(u)
	//	if err != nil {
	//		global.EVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
	//		response.FailWithMessage("用户名不存在或者密码错误", c)
	//		return
	//	}
	//	if user.Enable != 1 {
	//		global.EVA_LOG.Error("登陆失败! 用户被禁止登录!")
	//		response.FailWithMessage("用户被禁止登录", c)
	//		return
	//	}
	//	b.TokenNext(c, *user)
	//	return
	//}
}

// TokenNext 登录以后签发jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils2.JWT{SigningKey: []byte(global.EVA_CONFIG.JWT.SigningKey)} // 唯一签名

	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})

	token, err := j.CreateToken(claims)

	if err != nil {
		global.EVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	// 多点登录， 其他地点也能登录
	if !global.EVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			Token:     token,
			Username:  user.Username,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "多点登录成功", c)
		return
	}
	// 单点登录，登录后，jwt作废，其他地点登录会失败
	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.EVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			Token:     token,
			Username:  user.Username,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.EVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			Token:     token,
			Username:  user.Username,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// 发送短信
func (b *BaseApi) SendMessage(c *gin.Context) {
	mobile := c.Query("mobile")
	template := c.Query("template")
	var template_id string
	switch template {
	case "register":
		template_id = global.EVA_CONFIG.TencentSMS.Template.Register
	case "login":
		template_id = global.EVA_CONFIG.TencentSMS.Template.Login
	case "re_pwd":
		template_id = global.EVA_CONFIG.TencentSMS.Template.RePwd
	default:
		response.FailWithMessage("模版类型错误", c)
	}
	// 1. 判断手机号码是否在60秒内曾经发送过短信
	_, err := global.EVA_REDIS.Get(context.Background(), "mobile"+mobile).Result()
	if err != redis.Nil {
		response.FailWithMessage("短信已发送，请勿重复发送！", c)
		return
	}
	// 避免绕过前段，使用接口去重复请求发短信

	// 2. 生成短信验证码
	rand.Seed(time.Now().Unix()) //随机种子
	sms_code := rand.Intn(999999-100000) + 100000
	//随机生成100以内的正整数
	fmt.Println(sms_code)

	// 3. 保存短信验证码到redis[使用事务把多条命令集中发送给redis]
	// 创建管道对象
	pipe := global.EVA_REDIS.Pipeline()
	// 开启事务【无法管理数据库的读取数据操作】
	// 把事务中要完成的所有操作，写入到管道中
	pipe.Set(context.Background(), "register"+mobile, sms_code, 180*time.Second)
	pipe.Set(context.Background(), "mobile"+mobile, nil, 30*time.Second)
	// 执行事务
	_, err = pipe.Exec(context.Background())
	if err != nil {
		response.SFailWithMessage("服务端错误", c)
	}
	// 4. 调用短信sdk，发送短信
	//try:
	fmt.Println(mobile, template_id)
	//	ssl._create_default_https_context = ssl._create_unverified_context
	//	from mycelery.sms.tasks import send_sms
	//	from mycelery.mail.tasks import send_mail
	// send_mail.delay()
	// response = send_sms.delay(mobile, template_id, [sms_code, ])

	// ccp = CCP()
	// ret = ccp.send_template_sms(mobile, [sms_code, constants.SMS_EXPIRE_TIME//60], constants.SMS_TEMPLATE_ID)
	// if not ret:
	//     log.error("用户注册短信发送失败！手机号：%s" % mobile)
	//     return Response({"message":"发送短信失败！"},status=status.HTTP_500_INTERNAL_SERVER_ERROR)
	//except:
	//	return Response({"message":"发送短信失败！"},status=status.HTTP_500_INTERNAL_SERVER_ERROR)

	// 5. 响应发送短信的结果
	response.OkWithDetailed(map[string]interface{}{"sms_code": sms_code}, "发送成功", c)
}

// 手机号验证
func (b *BaseApi) CheckPhone(c *gin.Context) {
	phone := c.Param("phone")
	err := userService.CheckPhone(phone)
	if err != nil {
		response.SFailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Register
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /user/register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 重复密码校验
	if r.Password != r.RePassword {
		response.FailWithMessage("确认密码不一致", c)
	}
	// 验证码校验
	redis_Captcha, err := global.EVA_REDIS.Get(context.Background(), "register"+r.Phone).Result()
	if err != nil {
		response.FailWithMessage("验证码超时或错误", c)
		return
	}
	if err != nil {
		response.FailWithMessage("验证码类型错误", c)
		return
	}
	if redis_Captcha != r.Captcha {
		response.FailWithMessage("验证码超时或错误", c)
		return
	}
	//global.EVA_REDIS.Get(r.Phone)
	if len(r.AuthorityIds) > 0 {
		r.AuthorityId = r.AuthorityIds[0]
		for _, i := range r.AuthorityIds {
			if i < r.AuthorityId {
				r.AuthorityId = i
			} else {
				continue
			}
		}
	} else {
		r.AuthorityId = 888
		r.AuthorityIds = []uint{888}
	}

	err = utils2.Verify(r, utils2.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.EVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

// ChangePassword
// @Tags      SysUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router    /user/changePassword [post]
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils2.Verify(req, utils2.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils2.GetUserID(c)
	u := &system.SysUser{EVA_MODEL: global.EVA_MODEL{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.EVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetUserList
// @Tags      SysUser
// @Summary   分页获取用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /user/getUserList [post]
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils2.Verify(pageInfo, utils2.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.EVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// SetUserAuthority
// @Tags      SysUser
// @Summary   更改用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuth          true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /user/setUserAuthority [post]
func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if UserVerifyErr := utils2.Verify(sua, utils2.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils2.GetUserID(c)
	err = userService.SetUserAuthority(userID, sua.AuthorityId)
	if err != nil {
		global.EVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := utils2.GetUserInfo(c)
	j := &utils2.JWT{SigningKey: []byte(global.EVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims.AuthorityId = sua.AuthorityId
	if token, err := j.CreateToken(*claims); err != nil {
		global.EVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("new-token", token)
		c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
		response.OkWithMessage("修改成功", c)
	}
}

// SetUserAuthorities
// @Tags      SysUser
// @Summary   设置用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuthorities   true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /user/setUserAuthorities [post]
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.SetUserAuthorities(sua.ID, sua.AuthorityIds)
	if err != nil {
		global.EVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// DeleteUser
// @Tags      SysUser
// @Summary   删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "用户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除用户"
// @Router    /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils2.Verify(reqId, utils2.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils2.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		global.EVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// SetUserInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils2.Verify(user, utils2.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(user.AuthorityIds) != 0 {
		err = userService.SetUserAuthorities(user.ID, user.AuthorityIds)
		if err != nil {
			global.EVA_LOG.Error("设置失败!", zap.Error(err))
			response.FailWithMessage("设置失败", c)
			return
		}
	}
	err = userService.SetUserInfo(system.SysUser{
		EVA_MODEL: global.EVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.EVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// SetSelfInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/SetSelfInfo [put]
func (b *BaseApi) SetSelfInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils2.GetUserID(c)
	err = userService.SetUserInfo(system.SysUser{
		EVA_MODEL: global.EVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.EVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// GetUserInfo
// @Tags      SysUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /user/getUserInfo [get]
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils2.GetUserUuid(c)
	ReqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.EVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
}

// ResetPassword
// @Tags      SysUser
// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysUser                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "重置用户密码"
// @Router    /user/resetPassword [post]
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var user system.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(user.ID)
	if err != nil {
		global.EVA_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("重置成功", c)
}
