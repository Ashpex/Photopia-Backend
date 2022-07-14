package entity

type Comment struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Content   string `json:"content"`
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
