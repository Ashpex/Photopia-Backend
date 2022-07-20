package repository

import (
	"example.com/gallery/entity"
	"gorm.io/gorm"
)

type LikeRepository interface {
	Like(like entity.Like) entity.Like
	Unlike(like entity.Like)
	AllLikes(postID uint64) []entity.Like
	CountLikes(postID uint64) int
}

type likeConnection struct {
	connection *gorm.DB
}

func NewLikeRepository(databaseConnection *gorm.DB) LikeRepository {
	return &likeConnection{
		connection: databaseConnection,
	}
}

func (db *likeConnection) Like(like entity.Like) entity.Like {
	db.connection.Save(&like)
	db.connection.Preload("User").Find(&like)
	return like
}

func (db *likeConnection) Unlike(like entity.Like) {
	db.connection.Delete(&like).Where("post_id = ? AND user_id = ?", like.PostID, like.UserID)
	db.connection.Preload("User").Find(&like)
}

func (db *likeConnection) AllLikes(postID uint64) []entity.Like {
	var likes []entity.Like
	db.connection.Preload("User").Find(&likes, "post_id = ?", postID)
	return likes
}

func (db *likeConnection) CountLikes(postID uint64) int {
	var likes []entity.Like
	db.connection.Preload("Like").Find(&likes, "post_id = ?", postID)
	return len(likes)
}
