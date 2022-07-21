package repository

import (
	"example.com/gallery/entity"
	"gorm.io/gorm"
)

type FollowerRepository interface {
	Follow(follower entity.Follower) entity.Follower
	Unfollow(follower entity.Follower)
	AllFollower(userID uint64) []entity.Follower
	AllFollowing(userID uint64) []entity.Follower
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
	db.connection.Preload("User").Find(&follower)
	db.connection.Preload("Followers").Find(&follower)
	return follower
}

func (db *followerConnection) Unfollow(follower entity.Follower) {
	db.connection.Where("user_id = ? AND follower_id = ?", follower.UserID, follower.FollowerID).Delete(&follower)
	db.connection.Preload("User").Find(&follower)
}

func (db *followerConnection) AllFollower(userID uint64) []entity.Follower {
	var followers []entity.Follower
	db.connection.Preload("User").Find(&followers, "user_id = ?", userID)
	return followers
}

func (db *followerConnection) AllFollowing(userID uint64) []entity.Follower {
	var followers []entity.Follower
	db.connection.Preload("User").Find(&followers, "follower_id = ?", userID)
	return followers
}
