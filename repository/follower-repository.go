package repository

import (
	"example.com/gallery/entity"
	"gorm.io/gorm"
)

type FollowerRepository interface {
	Follow(follower entity.Follower) entity.Follower
	Unfollow(follower entity.Follower)
	AllFollower(userID uint64) []entity.Follower
}

type followerConnection struct {
	connection *gorm.DB
}

func NewFollowerRepository(databaseConnection *gorm.DB) FollowerRepository {
	return &followerConnection{
		connection: databaseConnection,
	}
}

func (db *followerConnection) Follow(follower entity.Follower) entity.Follower {
	db.connection.Save(&follower)
	return follower
}

func (db *followerConnection) Unfollow(follower entity.Follower) {
	db.connection.Delete(&follower)
}

func (db *followerConnection) AllFollower(userID uint64) []entity.Follower {
	var followers []entity.Follower
	db.connection.Preload("User").Find(&followers, "user_id = ?", userID)
	return followers
}
