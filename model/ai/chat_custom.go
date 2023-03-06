package ai

import "time"

type ChatCustom struct {
	User       string    `bson:"user" json:"user"`
	UserID     int64     `bson:"user_id" json:"user_id"`
	TotalToken int       `bson:"total_token" json:"total_token"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
}
