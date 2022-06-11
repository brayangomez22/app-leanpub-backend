package datastore

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	databse = "leanpub"
	users   = "users"
)

type MongoGatewayImpl struct {
	client *mongo.Client
}

func NewMongoGatewayImpl()