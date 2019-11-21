package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type GeolocationRepository struct {
	dbClient *mongo.Database
}

type LocationDocument struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID string `bson:"userID" json:"userId"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
	SharedIDs []string `bson:"sharedIDs" json:"sharedIDs"`
	Longitude float64
	Latitude float64
}