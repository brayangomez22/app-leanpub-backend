package datastore

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"leanpub-app/domain"
	"leanpub-app/domain/model"
	"os"
	"time"
)

const (
	database = "leanpub"
	users    = "users"
)

type MongoGatewayImpl struct {
	client *mongo.Client
}

func NewMongoGatewayImpl() domain.DatabaseGateway {
	return &MongoGatewayImpl{}
}

func (mongoImpl MongoGatewayImpl) Setup() {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	opt := options.Client()
	opt.ApplyURI(os.Getenv("mongo.url"))
	mongoImpl.client, err = mongo.Connect(ctx, opt)

	if err != nil {
		panic(err)
	}
}

func (mongoImpl MongoGatewayImpl) SaveUser(user *model.User) (*model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	opts := options.Update().SetUpsert(true)
	collection := mongoImpl.client.Database(database).Collection(users)

	id, _ := uuid.NewRandom()
	user.Id = id.String()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": user.Id}, bson.D{{"$set", user}}, opts)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (mongoImpl MongoGatewayImpl) ValidateUser(registeredUser *model.RegisteredUser, user *model.User) (*model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(users)

	err := collection.FindOne(ctx, bson.M{"email": registeredUser.Email}).Decode(&user)
	if err != nil {
		return nil, errors.New("INVALID_USER_OR_PASSWORD")
	}

	if user.Password != registeredUser.Password {
		return nil, errors.New("INVALID_USER_OR_PASSWORD")
	}

	return user, nil
}

func (mongoImpl MongoGatewayImpl) GetUsers() (*[]model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(users)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []model.User
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (mongoImpl MongoGatewayImpl) DeleteUser(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(users)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return err
}

func (mongoImpl MongoGatewayImpl) UpdateUser(user *model.User) (*model.User, error) {
	var userE *model.User
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	opts := options.Update().SetUpsert(true)
	collection := mongoImpl.client.Database(database).Collection(users)

	err := collection.FindOne(ctx, bson.M{"_id": user.Id}).Decode(&userE)
	if err != nil {
		return nil, errors.New("USER_NOT_FOUND")
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": user.Id}, bson.D{{"$set", user}}, opts)

	if err != nil {
		return nil, err
	}

	return user, nil
}
