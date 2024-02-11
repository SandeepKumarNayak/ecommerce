package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductName string             `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Rating      int                `json:"rating" bson:"rating"`
}
