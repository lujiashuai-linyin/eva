package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"eva/global"
	"eva/model/ai"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
	"time"
)

type ChatService struct{}

func (c *ChatService) Question(model string, content []ai.Message) (map[string]interface{}, error) {
	body := map[string]interface{}{}
	// 设置 OpenAI API 请求的参数
	reqBody := map[string]interface{}{
		"model":    model,
		"messages": content,
	}

	// 将请求参数转换为 JSON 格式
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return body, err
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqJSON))
	if err != nil {
		return body, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	// 读取响应数据并输出
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return body, err
	}
	//fmt.Printf("body=%+v", body)
	return body, nil
}

// 添加记录到mongo

func (c *ChatService) AddMessage(topic string, topicID int64, user string, user_id int64, message []ai.Message) error {
	// 先查询是否已有该条记录
	filter := bson.M{"topic": topic, "topic_id": topicID, "user_id": user_id}
	chat_table := global.EVA_MongoDB.Database("chat_app").Collection("chat_history")
	result := chat_table.FindOne(context.Background(), filter)

	chat := ai.Chat{
		Topic:     topic,
		TopicID:   topicID,
		User:      user,
		UserID:    user_id,
		Message:   message,
		CreatedAt: time.Now(),
	}

	// 如果查询到了记录，则更新 Message 字段
	if result.Err() == nil {
		fmt.Println("查到了")
		update := bson.M{"$set": chat}
		_, err := chat_table.UpdateOne(context.Background(), filter, update)
		return err
	}

	// 如果查询不到记录，则插入新记录
	_, err := chat_table.InsertOne(context.Background(), chat)
	return err
}

func (c *ChatService) GetTopicMessage(topic_id int, topic string, user_id int) ([]ai.Message, error) {
	filter := bson.M{"topic": topic, "topic_id": topic_id, "user_id": user_id}
	chat_table := global.EVA_MongoDB.Database("chat_app").Collection("chat_history")
	result := chat_table.FindOne(context.Background(), filter)
	if result.Err() != nil {
		fmt.Println(result.Err())
		return nil, nil
	}
	var chats ai.Chat
	if err := result.Decode(&chats); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return chats.Message, nil
}

func (c *ChatService) GetTopicList(user_id int) ([]ai.TopicList, error) {
	filter := bson.M{"user_id": user_id}
	chat_table := global.EVA_MongoDB.Database("chat_app").Collection("chat_history")
	result, err := chat_table.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(result.Err())
		return nil, nil
	}
	defer result.Close(context.Background())
	var topics []ai.TopicList
	if err = result.All(context.Background(), &topics); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("topics:%+v", topics)
	return topics, nil
}

func (c *ChatService) AddToken(user string, user_id int64, token_number int) error {
	// 先查询是否已有该条记录
	filter := bson.M{"user_id": user_id}
	chat_table := global.EVA_MongoDB.Database("chat_app").Collection("chat_custom")
	result := chat_table.FindOne(context.Background(), filter)

	// 如果查询到了记录，则更新 Message 字段
	if result.Err() == nil {
		var chatUpdate ai.ChatCustom
		err := result.Decode(&chatUpdate)
		//fmt.Println(chatUpdate.TotalToken)
		chatUpdate.TotalToken += token_number
		update := bson.M{"$set": chatUpdate}
		_, err = chat_table.UpdateOne(context.Background(), filter, update)
		return err
	} else {
		chat := ai.ChatCustom{
			User:       user,
			UserID:     user_id,
			TotalToken: token_number,
			CreatedAt:  time.Now(),
		}

		// 如果查询不到记录，则插入新记录
		_, err := chat_table.InsertOne(context.Background(), chat)
		return err
	}
}
