package dto

// LoginDTO is used when client post from /login endpoint
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required" validate:"min=8"`
}
