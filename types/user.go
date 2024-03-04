package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	UID            string             `bson:"uid"`
	FirstName      string             `bson:"first_name"`
	LastName       string             `bson:"last_name"`
	Email          string             `bson:"email"`
	PhoneNumber    string             `bson:"phone_number"`
	Password       string             `bson:"password"`
	IdentityNumber string             `bson:"identity_number"`
	CreatedAt      time.Time          `bson:"created_at"`
}

func (u *User) Response() *ResponseUser {
	return &ResponseUser{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		CreatedAt:   u.CreatedAt,
	}
}

type ResponseUser struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type SignUpData struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PhoneNumber    string `json:"phone_number"`
	IdentityNumber string `json:"identity_number"`
}

type SignInData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
