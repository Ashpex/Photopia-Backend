package dto

type UnfollowDTO struct {
	UserID     string `json:"user_id" form:"user_id" binding:"required"`
	FollowerID string `json:"follower_id" form:"follower_id" binding:"required"`
}

type FollowDTO struct {
	UserID     string `json:"user_id" form:"user_id" binding:"required"`
	FollowerID string `json:"follower_id" form:"follower_id" binding:"required"`
}
