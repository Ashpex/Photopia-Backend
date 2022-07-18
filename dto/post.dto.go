package dto

//PostUpdateDTO is a model that client use when updating a post
type PostUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	//PhotoUrl    string   `json:"photo_url" form:"photo_url" binding:"required"`
	//Photo  *multipart.FileHeader `json:"photo" form:"photo"`
	//TopicID []uint64 `json:"topic_id" form:"topic_id" binding:"required"`

}

// PostCreateDTO is used by client when POST create new post
type PostCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	//TopicID     []uint64 `json:"topic_id" form:"topic_id" binding:"required"`
	//PhotoUrl    string   `json:"photo_url" form:"photo_url" binding:"required"`
	//Photo  *multipart.FileHeader `json:"photo" form:"photo"`
}
