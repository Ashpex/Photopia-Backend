package entity

// Post struct represents a post in the database.
type Post struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"json:"user"`
	PhotoUrl    string `json:"photo_url"`
	comment     []Comment
	topic       uint64 `gorm:"not null" json:"-"`
	Topic       Topic  `gorm:"foreignkey:TopicID;constraint:onUpdate:CASCADE,onDelete:CASCADE"json:"topic"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
