package repository

import (
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/entity"
	"gorm.io/gorm"
	"log"
)

type TopicRepository interface {
	InsertTopic(topic entity.Topic) entity.Topic
	UpdateTopic(topic entity.Topic) entity.Topic
	DeleteTopic(topic entity.Topic)
	AllTopic() []entity.Topic
	FindTopicByID(topicID uint64) entity.Topic
}

type topicConnection struct {
	connection *gorm.DB
}

func NewTopicRepository(databaseConnection *gorm.DB) TopicRepository {
	return &topicConnection{
		connection: databaseConnection,
	}
}

func (db *topicConnection) InsertTopic(topic entity.Topic) entity.Topic {
	err := db.connection.Save(&topic)
	if err != nil {
		log.Println(err)
	}
	return topic
}

func (db *topicConnection) UpdateTopic(topic entity.Topic) entity.Topic {
	db.connection.Save(&topic)
	return topic
}

func (db *topicConnection) DeleteTopic(topic entity.Topic) {
	err := db.connection.Delete(&topic)
	if err != nil {
		log.Println(err)
	}
}

func (db *topicConnection) AllTopic() []entity.Topic {
	var topics []entity.Topic
	db.connection.Find(&topics)
	return topics
}

func (db *topicConnection) FindTopicByID(topicID uint64) entity.Topic {
	var topic entity.Topic
	db.connection.Find(&topic, topicID)
	return topic
}
