package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName       string             `json:"first_name" bson:"first_name,omitempty"`
	LastName        string             `json:"last_name" bson:"last_name,omitempty"`
	Email           string             `json:"email" bson:"email,omitempty"`
	Password        string             `json:"password" bson:"password,omitempty"`
	ConfirmPassword string             `json:"confirm_password" bson:"confirm_password,omitempty"`
}
