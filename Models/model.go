package model

import (
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Show struct{
	Id primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,o,itempty"`
	Name string 			`json:"name,omitempty"`
	Watched bool 			`json:"watched,omitempty"`
}