package dto

// CommentCreateDTO is used by client when POST create new comment
type CommentCreateDTO struct {
	ID      uint64 `json:"id" form:"id" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	UserID  uint64 `json:"user_id" form:"user_id" binding:"required"`
	PostID  uint64 `json:"post_id" form:"post_id" binding:"required"`
}

// CommentUpdateDTO is used by client when PUT update comment
type CommentUpdateDTO struct {
	ID      uint64 `json:"id" form:"id" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	UserID  uint64 `json:"user_id" form:"user_id" binding:"required"`
	PostID  uint64 `json:"post_id" form:"post_id" binding:"required"`
}
