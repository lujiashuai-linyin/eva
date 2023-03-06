package response

import "eva/model/ai"

type ChatResp struct {
	Code int `json:"code"`
	Data struct {
		Choices []struct {
			FinishReason string     `json:"finish_reason"`
			Index        int        `json:"index"`
			Message      ai.Message `json:"message"`
		} `json:"choices"`
		Created int64  `json:"created"`
		ID      string `json:"id"`
		Model   string `json:"model"`
		Object  string `json:"object"`
		Usage   struct {
			CompletionTokens int `json:"completion_tokens"`
			PromptTokens     int `json:"prompt_tokens"`
			TotalTokens      int `json:"total_tokens"`
		} `json:"usage"`
	} `json:"data"`
	Msg string `json:"msg"`
}
