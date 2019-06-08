package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Post struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
	Created time.Time          `json:"created,omitempty" bson:"created,omitempty"`
}

type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Surname        string             `json:"surname,omitempty" bson:"surname,omitempty"`
	Email          string             `json:"email,omitempty" bson:"email,omitempty"`
	HashedPassword []byte             `json:"hashedPassword,omitempty" bson:"hashedPassword,omitempty"`
	Created        time.Time          `json:"created,omitempty" bson:"created,omitempty"`
}
