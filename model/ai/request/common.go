package request

import "eva/model/ai"

// 模型请求体
type QuestionReq struct {
	Model   string       `json:"model"`
	Topic   string       `json:"topic"`
	TopicID int          `json:"topic_id"`
	Content []ai.Message `json:"content"`
}
