package entity

// Post struct represents a post in the config.

type Post struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	//PhotoUrl    string `json:"photo_url"`
	//Photo  *multipart.FileHeader `json:"photo" form:"photo" binding:"required"`
}
