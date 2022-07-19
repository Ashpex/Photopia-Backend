package repository

import (
	"example.com/gallery/entity"
	"gorm.io/gorm"
)

type CommentRepository interface {
	InsertComment(comment entity.Comment) entity.Comment
	UpdateComment(comment entity.Comment) entity.Comment
	DeleteComment(comment entity.Comment)
	AllComment() []entity.Comment
	FindCommentByID(commentID uint64) entity.Comment
	FindCommentByPostID(postID uint64) []entity.Comment
}

type commentConnection struct {
	connection *gorm.DB
}

func NewCommentRepository(databaseConnection *gorm.DB) CommentRepository {
	return &commentConnection{
		connection: databaseConnection,
	}
}

func (db *commentConnection) InsertComment(comment entity.Comment) entity.Comment {
	db.connection.Save(&comment)
	db.connection.Preload("User").Find(&comment)
	db.connection.Preload("Post").Find(&comment)
	return comment
}

func (db *commentConnection) UpdateComment(comment entity.Comment) entity.Comment {
	db.connection.Save(&comment)
	db.connection.Preload("User").Find(&comment)
	db.connection.Preload("Post").Find(&comment)
	return comment
}

func (db *commentConnection) DeleteComment(comment entity.Comment) {
	db.connection.Delete(&comment)
}

func (db *commentConnection) AllComment() []entity.Comment {
	var comments []entity.Comment
	db.connection.Preload("User").Find(&comments)
	db.connection.Preload("Post").Find(&comments)
	return comments
}

func (db *commentConnection) FindCommentByID(commentID uint64) entity.Comment {
	var comment entity.Comment
	db.connection.Preload("Post").Find(&comment, commentID)
	return comment
}

func (db *commentConnection) FindCommentByPostID(postID uint64) []entity.Comment {
	var comments []entity.Comment
	db.connection.Preload("Post").Find(&comments, "comment_id = ?", postID)
	return comments
}
