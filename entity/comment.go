package entity

type Comment struct {
	ID      uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Content string `json:"content"`
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
}
