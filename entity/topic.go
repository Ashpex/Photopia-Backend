package entity

type Topic struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
