package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`    // bson:"_id,omitempty" is used to omit the field when the value is empty
	Name string             `json:"name" bson:"name,omitempty"` // bson:"name,omitempty" is used to omit the field when the value is empty
}
