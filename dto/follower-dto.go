package dto

type UnfollowDTO struct {
	TargetUserID uint64 `json:"follower_id" form:"follower_id" binding:"required"`
}

type FollowDTO struct {
	TargetUserID uint64 `json:"user_id" form:"user_id" binding:"required"`
}
