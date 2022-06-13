package models

import "time"

type SocialNetwork struct {
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}

type User struct {
	Id              string          `json:"id" bson:"_id"`
	Name            string          `json:"name" bson:"name"`
	Password        string          `json:"password" bson:"password"`
	Email           string          `json:"email" bson:"email"`
	About           string          `json:"about" bson:"about"`
	AvatarUrl       string          `json:"avatarUrl" bson:"avatarUrl"`
	HasSubscription bool            `json:"hasSubscription" bson:"hasSubscription"`
	IsAuthor        bool            `json:"isAuthor" bson:"isAuthor"`
	IsAdmin         bool            `json:"isAdmin" bson:"isAdmin"`
	CreatedAt       time.Time       `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt" bson:"updatedAt"`
	SocialNetworks  []SocialNetwork `json:"socialNetworks" bson:"socialNetworks"`
}

type RegisteredUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
