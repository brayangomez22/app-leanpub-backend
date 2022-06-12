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
	database    = "leanpub"
	users       = "users"
	books       = "books"
	bookContent = "bookContent"
)

type MongoGatewayImpl struct {
	client *mongo.Client
}

func NewMongoGatewayImpl() domain.DatabaseGateway {
	return &MongoGatewayImpl{}
}

func (mongoImpl *MongoGatewayImpl) Setup() {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	opt := options.Client()
	opt.ApplyURI(os.Getenv("mongo.url"))
	mongoImpl.client, err = mongo.Connect(ctx, opt)

	if err != nil {
		panic(err)
	}
}

func (mongoImpl *MongoGatewayImpl) SaveUser(user *model.User) (*model.User, error) {
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

func (mongoImpl *MongoGatewayImpl) ValidateUser(registeredUser *model.RegisteredUser, user *model.User) (*model.User, error) {
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

func (mongoImpl *MongoGatewayImpl) GetUsers() (*[]model.User, error) {
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

func (mongoImpl *MongoGatewayImpl) GetUserById(id string) (*model.User, error) {
	var user *model.User
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(users)

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, errors.New("USER_NOT_FOUND")
	}

	return user, nil
}

func (mongoImpl *MongoGatewayImpl) DeleteUser(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(users)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return err
}

func (mongoImpl *MongoGatewayImpl) UpdateUser(user *model.User) (*model.User, error) {
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

func (mongoImpl *MongoGatewayImpl) SaveBook(book *model.Book) (*model.Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	opts := options.Update().SetUpsert(true)
	bookCollection := mongoImpl.client.Database(database).Collection(books)

	id, _ := uuid.NewRandom()
	book.Id = id.String()

	_, err := bookCollection.UpdateOne(ctx, bson.M{"_id": book.Id}, bson.D{{"$set", book}}, opts)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (mongoImpl *MongoGatewayImpl) GetBooks() (*[]model.Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var books []model.Book
	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, err
	}

	return &books, nil
}

func (mongoImpl *MongoGatewayImpl) GetBookById(id string) (*model.Book, error) {
	var book *model.Book
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return nil, errors.New("BOOK_NOT_FOUND")
	}

	return book, nil
}

func (mongoImpl *MongoGatewayImpl) GetBookByAuthor(authorId string) (*model.Book, error) {
	var book *model.Book
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	err := collection.FindOne(ctx, bson.M{"authors.authorId": authorId}).Decode(&book)
	if err != nil {
		return nil, errors.New("BOOK_NOT_FOUND")
	}

	return book, nil
}

func (mongoImpl *MongoGatewayImpl) GetBookByCategory(category string) (*model.Book, error) {
	var book *model.Book
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	err := collection.FindOne(ctx, bson.D{
		{"categories",
			bson.D{
				{"$all",
					bson.A{
						category,
					},
				},
			},
		},
	})

	if err != nil {
		return nil, errors.New("BOOK_NOT_FOUND")
	}

	return book, nil
}

func (mongoImpl *MongoGatewayImpl) DeleteBook(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return err
}

func (mongoImpl *MongoGatewayImpl) UpdateBook(book *model.Book) (*model.Book, error) {
	var bookE *model.Book
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	opts := options.Update().SetUpsert(true)
	collection := mongoImpl.client.Database(database).Collection(books)

	err := collection.FindOne(ctx, bson.M{"_id": book.Id}).Decode(&bookE)
	if err != nil {
		return nil, errors.New("BOOK_NOT_FOUND")
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": book.Id}, bson.D{{"$set", book}}, opts)
	if err != nil {
		return nil, err
	}

	return book, nil
}
