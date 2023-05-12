package ai

import (
	"eva/model/ai"
	AIReq "eva/model/ai/request"
	"eva/model/common/response"
	"eva/model/system/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ChatApi struct{}

func (b *ChatApi) DeleteTopic(c *gin.Context) {
	claims, _ := c.Get("claims")
	userID := claims.(*request.CustomClaims).ID
	res := map[string]interface{}{}
	err := c.ShouldBindJSON(&res)
	if err != nil {
		response.FailWithMessage("参数错误！", c)
		return
	}
	id := res["id"].(float64)
	err = chatService.Delete(int(id), int(userID))
	if err != nil {
		fmt.Println(err)
		response.SFailWithMessage("fail！", c)
		return
	}
	response.OkWithMessage("success", c)
}

func (b *ChatApi) Question(c *gin.Context) {
	// 获取用户信息
	var req AIReq.QuestionReq
	claims, _ := c.Get("claims")
	userID := claims.(*request.CustomClaims).ID
	user := claims.(*request.CustomClaims).NickName
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("参数错误！", c)
		return
	}

	// 请求内容，gpt3.5引擎，返回恢复
	data, err := chatService.Question(req.Model, req.Content)
	if err != nil {
		fmt.Println(err)
		response.SFailWithMessage("fail！", c)
		return
	}

	// 更新mongodb消息
	choices := data["choices"].([]interface{})
	usage := data["usage"].(map[string]interface{})
	total_tokens := usage["total_tokens"].(float64)
	messages := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	content := messages["content"].(string)
	resp := ai.Message{Role: "assistant", Content: content}
	message := append(req.Content, resp)
	//fmt.Println("totle_token", total_tokens)
	go func(total_tokens int) {
		err = chatService.AddToken(user, int64(userID), total_tokens)
	}(int(total_tokens))
	fmt.Printf("topic_id:%d", req.TopicID)
	err = chatService.AddMessage(req.Topic, int64(req.TopicID), user, int64(userID), message)
	if err != nil {
		fmt.Println(err)
		response.SFailWithMessage("fail！", c)
		return
	}

	response.OkWithDetailed(data, "success", c)
}

func (b *ChatApi) GetTopicMessage(c *gin.Context) {
	topic_id := c.Query("topic_id")
	topicID, err := strconv.Atoi(topic_id)
	if err != nil {
		fmt.Println(err)
		response.SFailWithMessage("fail！", c)
		return
	}
	topic := c.Query("topic")
	claims, _ := c.Get("claims")
	userID := claims.(*request.CustomClaims).ID

	// 请求特定topic内容
	data, err := chatService.GetTopicMessage(topicID, topic, int(userID))
	if err != nil {
		response.SFailWithMessage("fail！", c)
		return
	}
	response.OkWithDetailed(data, "success", c)
}

func (b *ChatApi) GetTopicList(c *gin.Context) {
	claims, _ := c.Get("claims")
	userID := claims.(*request.CustomClaims).ID
	data, err := chatService.GetTopicList(int(userID))
	if err != nil {
		response.SFailWithMessage("fail！", c)
		return
	}
	response.OkWithDetailed(data, "success", c)
}
