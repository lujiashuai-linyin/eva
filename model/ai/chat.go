package ai

import "time"

type Message struct {
	Role    string `json:"role" bson:"role"`
	Content string `json:"content" bson:"content"`
}

type Chat struct {
	Topic     string    `bson:"topic"`
	TopicID   int64     `bson:"topic_id"`
	Message   []Message `bson:"message"`
	User      string    `bson:"user"`
	UserID    int64     `bson:"user_id"`
	CreatedAt time.Time `bson:"created_at"`
}

type TopicList struct {
	TopicID int64  `bson:"topic_id" json:"id"`
	Topic   string `bson:"topic" json:"topic"`
}
