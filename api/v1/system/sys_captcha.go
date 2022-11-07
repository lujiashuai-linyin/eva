package system

import (
	"eva/global"
	"eva/model/common/response"
	systemRes "eva/model/system/response"
	"eva/utils/gt3"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"strconv"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
//var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router    /base/captcha [post]
// height 高度 png 像素高度
// width  宽度 png 像素高度
// length 验证码默认位数
// maxSkew 单个数字的最大绝对倾斜因子
// dotCount 背景圆圈的数量
func (b *BaseApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.EVA_CONFIG.Captcha.ImgHeight, global.EVA_CONFIG.Captcha.ImgWidth, global.EVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global.EVA_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.EVA_CONFIG.Captcha.KeyLong,
	}, "验证码获取成功", c)
}

func (b *BaseApi) GetGtCaptcha(c *gin.Context) {
	username := c.Query("username")
	fmt.Println("username", username)
	if username == "" {
		global.EVA_LOG.Error("请求参数错误!")
		response.FailWithMessage("请求参数错误！", c)
		return
	}
	user, err := userService.FindUserByUserName(username)
	if err != nil {
		global.EVA_LOG.Error("无当前用户!", zap.Error(err))
		response.FailWithMessage("无当前用户！", c)
		return
	}
	gt := gt3.NewClient(global.EVA_CONFIG.Captcha.GeeTestID, global.EVA_CONFIG.Captcha.GeeTestKey)
	_, challenge := gt.PreProcess(int(user.ID))
	//if err != nil || res.Challenge == "0" {
	//	global.EVA_LOG.Error("验证码获取初始化错误!", zap.Error(err))
	//	response.SFailWithMessage("验证码获取初始化错误！", c)
	//	return
	//}
	fmt.Printf("registerResponse: %+v\n", challenge)
	//global.EVA_REDIS.Set(context.Background(), res.Challenge, user.ID, 60*time.Second)

	response.OkWithDetailed(map[string]interface{}{"challenge": challenge, "gt": global.EVA_CONFIG.Captcha.GeeTestID}, "获取验证码成功", c)
}

func (b *BaseApi) ValidGtCaptcha(c *gin.Context) {
	validate := map[string]interface{}{}
	err := c.ShouldBindJSON(&validate)
	if err != nil {
		response.FailWithMessage("参数错误。", c)
		return
	}
	user, err := userService.FindUserByUserName(validate["username"].(string))
	if err != nil {
		global.EVA_LOG.Error("无当前用户!", zap.Error(err))
		response.FailWithMessage("无当前用户！", c)
		return
	}
	gt := gt3.NewClient(global.EVA_CONFIG.Captcha.GeeTestID, global.EVA_CONFIG.Captcha.GeeTestKey)
	//userID := global.EVA_REDIS.Get(context.Background(), validate["geetest_challenge"].(string))
	//if userID == nil {
	//	response.FailWithMessage("验证码错误或已失效", c)
	//	return
	//}
	_, err = gt.Validate(validate["geetest_challenge"].(string), validate["geetest_seccode"].(string), strconv.Itoa(int(user.ID)))
	if err != nil {
		fmt.Printf("ValidateResponse Err: %v\n", err)
		response.SFailWithMessage("ValidateResponse！Err!", c)
		return
	}
	response.Ok(c)
}
