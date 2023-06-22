package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookSchema struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}
