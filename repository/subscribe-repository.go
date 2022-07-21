package repository

import (
	"example.com/gallery/entity"
	"gorm.io/gorm"
)

type SubscribeRepository interface {
	Subscribe(subscribe entity.Subscribe) entity.Subscribe
	Unsubscribe(subscribe entity.Subscribe)
	AllSubscribes(topicID uint64) []entity.Subscribe
	CountSubscribes(topicID uint64) int
}

type subscribeConnection struct {
	connection *gorm.DB
}

func NewSubscribeRepository(databaseConnection *gorm.DB) SubscribeRepository {
	return &subscribeConnection{
		connection: databaseConnection,
	}
}

func (db *subscribeConnection) Subscribe(subscribe entity.Subscribe) entity.Subscribe {
	db.connection.Save(&subscribe)
	db.connection.Preload("User").Find(&subscribe)
	return subscribe
}

func (db *subscribeConnection) Unsubscribe(subscribe entity.Subscribe) {
	db.connection.Delete(&subscribe).Where("topic_id = ? AND user_id = ?", subscribe.TopicID, subscribe.UserID)
	db.connection.Preload("User").Find(&subscribe)
}

func (db *subscribeConnection) AllSubscribes(topicID uint64) []entity.Subscribe {
	var subscribes []entity.Subscribe
	db.connection.Preload("User").Find(&subscribes, "topic_id = ?", topicID)
	return subscribes
}

func (db *subscribeConnection) CountSubscribes(topicID uint64) int {
	var subscribes []entity.Subscribe
	db.connection.Preload("Subscribe").Find(&subscribes, "topic_id = ?", topicID)
	return len(subscribes)
}
