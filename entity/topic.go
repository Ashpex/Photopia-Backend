package entity

type Topic struct {
	ID   uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"title"`
	//Posts *[]Post `gorm:"many2many:post_topic;association_join_table_foreignkey:topic_id;foreignkey:id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"posts,omitempty"`
	Posts *[]Post `json:"posts,omitempty"`
}
