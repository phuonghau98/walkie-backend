package main

import (
	"context"
	"github.com/phuonghau98/walkie3/srv/message/proto/message"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MessageDocument struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Sender string `bson:"sender" json:"sender"`
	Receiver string `bson:"receiver" json:"receiver"`
	Content string `bson:"content" json:"content"`
	IsRead bool `bson:"isRead" json:"isRead"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
//Tuyen: b4ce6937-6498-4275-8864-ab795c23560b
// Yen: d1fbfe17-fb3f-4433-8394-7157503eb130
type MessageRepository struct {
	dbClient *mongo.Database
}

func (repo *MessageRepository) GetMessages (userID string, milestone int64) []*MessageDocument{
	var messages []*MessageDocument
	conditions := bson.D{{"$or", bson.A{
		bson.D{{"sender", userID}},
		bson.D{{"receiver", userID}},
	}}}
	if milestone != 0 {
		conditions = append(conditions, bson.E{Key: "createdAt", Value: bson.D{{"$gt", time.Unix(milestone / 1000, 0)}}})
	}
	cur, err := repo.dbClient.Collection("message").Find(context.TODO(), conditions)
	if err != nil {
		return nil
	}
	for cur.Next(context.TODO()) {
		var tmpMessage MessageDocument
		_ = cur.Decode(&tmpMessage)
		messages = append(messages, &tmpMessage)
	}
	return messages
}

func (repo *MessageRepository) InsertMessage (message *message.Message) error {
	msg := MessageDocument{
		Sender:   message.Sender,
		Receiver: message.Receiver,
		Content:  message.Content,
		IsRead:   false,
		CreatedAt: time.Now(),
	}
	_, err := repo.dbClient.Collection("message").InsertOne(context.Background(), msg)
	return err
}

func (repo *MessageRepository) GetSentMessages (senderID string, receiverID string, limited int) ([]*MessageDocument, error) {
	var messages []*MessageDocument
	conditions := bson.D{{"$or", bson.A{
		bson.D{
		{"sender", senderID},
		{"receiver", receiverID},
		},
		bson.D{
		{"receiver", senderID},
		{"sender", receiverID},
		},
	}}}
	options := options2.Find()
	options.SetSort(map[string]int{"createdAt": -1})
	//if limited != -1 {
	//	options.SetLimit(10)
	//}
	cur, err := repo.dbClient.Collection("message").Find(context.TODO(), conditions, options)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var tmpMsg MessageDocument
		_ = cur.Decode(&tmpMsg)
		messages = append(messages, &tmpMsg)
	}
	return messages, nil
}