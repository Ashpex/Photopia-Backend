package entity

type Follower struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID     uint64 `gorm:"not null" json:"-"`
	FollowerID uint64 `gorm:"not null" json:"-"`
	//User   User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"json:"user"`
}