package repository

import (
	"example.com/gallery/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	InsertPost(post entity.Post) entity.Post
	UpdatePost(post entity.Post) entity.Post
	DeletePost(post entity.Post)
	AllPost() []entity.Post
	FindPostByID(ID uint64) entity.Post
}

type postConnection struct {
	connection *gorm.DB
}

// NewPostRepository is creates a new instance of PostRepository
func NewPostRepository(db *gorm.DB) PostRepository {
	return &postConnection{
		connection: db,
	}
}

func (db *postConnection) InsertPost(post entity.Post) entity.Post {
	db.connection.Save(&post)
	db.connection.Preload("User").Find(&post)
	return post
}

func (db *postConnection) UpdatePost(post entity.Post) entity.Post {
	db.connection.Save(&post)
	db.connection.Preload("User").Find(&post)
	return post
}

func (db *postConnection) DeletePost(post entity.Post) {
	db.connection.Delete(&post)
}
func (db *postConnection) AllPost() []entity.Post {
	var posts []entity.Post
	db.connection.Preload("User").Find(&posts)
	return posts
}

func (db *postConnection) FindPostByID(ID uint64) entity.Post {
	var post entity.Post
	db.connection.Preload("User").Find(&post, ID)
	return post
}
