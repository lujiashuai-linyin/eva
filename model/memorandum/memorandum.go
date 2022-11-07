package memorandum

import "eva/global"

type Memorandum struct {
	global.EVA_MODEL
	MemorandumTopicID int
	MemorandumTopic   MemorandumTopic
	TopicName         string `json:"topic_name"  form:"topic_name" gorm:"column:topic_name;"`
	Content           string `json:"content" form:"content" gorm:"column:content;type:text;comment:content"`
	Extra             string `json:"extra" form:"extra" gorm:"column:extra;comment:extra"`
}
