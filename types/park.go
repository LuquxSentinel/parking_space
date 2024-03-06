package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Park struct {
	ID           primitive.ObjectID `bson:"_id"`
	ParkID       string             `bson:"park_id"`
	Location     Location           `bson:"location"`
	Name         string             `bson:"name"`
	Images       []string           `bson:"images"`
	PricePerHour float64            `bson:"price_per_hour"`
	MaxCapacity  uint16             `bson:"max_capacity"`
	FreeSpace    uint16             `bson:"free_space"`
}

type Location struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}
