package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	EmployeeID string             `bson:"employee_id"`
	Password   string             `bson:"password"`
	Email      string             `bson:"email"`
}
