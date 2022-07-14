package dto

// BookUpdateDTO is used by client when updating a book (PUT)
type PostUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id" form:"user_id" binding:"required"`
	TopicID     uint64 `json:"topic_id" form:"topic_id" binding:"required"`
	PhotoUrl    string `json:"photo_url" form:"photo_url" binding:"required"`
}

// PostCreateDTO is used by client when POST create new post
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id" form:"user_id" binding:"required"`
	TopicID     uint64 `json:"topic_id" form:"topic_id" binding:"required"`
	PhotoUrl    string `json:"photo_url" form:"photo_url" binding:"required"`
}
