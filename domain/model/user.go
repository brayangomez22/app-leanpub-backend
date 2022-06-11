package model

import "time"

type SocialNetwork struct {
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}

type User struct {
	Id              string          `json:"_id" bson:"_id"`
	Name            string          `json:"name" bson:"name"`
	Email           string          `json:"email" bson:"email"`
	About           string          `json:"about" bson:"about"`
	AvatarUrl       string          `json:"avatarUrl" bson:"avatarUrl"`
	HasSubscription string          `json:"hasSubscription" bson:"hasSubscription"`
	IsAuthor        bool            `json:"isAuthor" bson:"isAuthor"`
	IsAdmin         bool            `json:"isAdmin" bson:"isAdmin"`
	CreatedAt       time.Time       `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt" bson:"updatedAt"`
	SocialNetworks  []SocialNetwork `json:"socialNetworks" bson:"socialNetworks"`
}
