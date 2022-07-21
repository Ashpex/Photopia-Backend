package entity

type Subscribe struct {
	ID      uint64 `gorm:"primary_key:auto_increment"`
	UserID  uint64 `gorm:"not null;uniqueIndex:idx_userid_topicid" json:"user_id" form:"user_id"`
	TopicID uint64 `gorm:"not null;uniqueIndex:idx_userid_topicid" json:"topic_id" form:"topic_id" binding:"required"`
}
