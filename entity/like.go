package entity

type Like struct {
	ID     uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID uint64 `gorm:"not null" json:"user_id"`
	PostID uint64 `gorm:"not null" json:"post_id"`
	User   User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
