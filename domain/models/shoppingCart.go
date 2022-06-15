package models

import "time"

type BookId struct {
	Book string `json:"book" bson:"book"`
}

type ShoppingCart struct {
	Id        string    `json:"id" bson:"_id"`
	UserId    string    `json:"userId" bson:"userId"`
	Books     []BookId  `json:"books" bson:"books"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
