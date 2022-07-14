package entity

// User struct represents a user in the database.
type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name" binding:"required"`
	Email     string `gorm:"uniqueIndex;type:varchar(255)" json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Token     string `gorm:"-" json:"token,omitempty"`
}
