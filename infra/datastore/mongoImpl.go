package datastore

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"leanpub-app/domain"
	"leanpub-app/domain/models"
	"os"
	"time"
)

const (
	database    = "leanpub"
	users       = "users"
	books       = "books"
	bookSections = "bookSections"
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

func (mongoImpl *MongoGatewayImpl) SaveUser(user *models.User) (*models.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	opts := options.Update().SetUpsert(true)
	collection := mongoImpl.client.Database(database).Collection(users)

	id, _ := uuid.NewRandom()
	user.Id = id.String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": user.Id}, bson.D{{"$set", user}}, opts)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (mongoImpl *MongoGatewayImpl) ValidateUser(registeredUser *models.RegisteredUser, user *models.User) (*models.User, error) {
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

func (mongoImpl *MongoGatewayImpl) GetUsers() (*[]models.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(users)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (mongoImpl *MongoGatewayImpl) GetUserById(id string) (*models.User, error) {
	var user *models.User
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

func (mongoImpl *MongoGatewayImpl) UpdateUser(user *models.User) (*models.User, error) {
	var userE *models.User
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	opts := options.Update().SetUpsert(true)
	collection := mongoImpl.client.Database(database).Collection(users)

	err := collection.FindOne(ctx, bson.M{"_id": user.Id}).Decode(&userE)
	if err != nil {
		return nil, errors.New("USER_NOT_FOUND")
	}

	user.UpdatedAt = time.Now()

	_, err = collection.UpdateOne(ctx, bson.M{"_id": user.Id}, bson.D{{"$set", user}}, opts)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (mongoImpl *MongoGatewayImpl) SaveBook(book *models.Book) (*models.Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	opts := options.Update().SetUpsert(true)
	collection := mongoImpl.client.Database(database).Collection(books)

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": book.Id}, bson.D{{"$set", book}}, opts)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (mongoImpl *MongoGatewayImpl) SaveBookSection(bookSection *models.BookSection) error {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(bookSections)

	_, err := collection.InsertOne(ctx, bookSection)
	return err
}

func (mongoImpl *MongoGatewayImpl) SaveBookSections(sections []interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(bookSections)

	_, err := collection.InsertMany(ctx, sections)
	return err
}

func (mongoImpl *MongoGatewayImpl) GetBooks() (*[]models.Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var books []models.Book
	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, err
	}

	return &books, nil
}

func (mongoImpl *MongoGatewayImpl) GetBookIndex(id string) (*[]models.BookContent, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	pipeline := make([]bson.D, 0, 0)
	queryPipeline := make([]bson.D, 0, 0)
	pipeline = append(pipeline, bson.D{{"$match", bson.D{{"_id", id}}}})
	pipeline = append(pipeline, bson.D{
		{"$project",
			bson.D{
				{"content.chapter", 1},
				{"content.sections", 1},
				{"_id", 0},
			},
		},
	})
	queryPipeline = append(queryPipeline, pipeline...)

	cursor, err := collection.Aggregate(ctx, queryPipeline)
	if err != nil {
		return nil, err
	}

	var response []models.BookContent
	var bookIndex []models.BookIndex
	err = cursor.All(ctx, &bookIndex)
	if err != nil {
		return nil, err
	}

	for _, book := range bookIndex[0].Content {
		response = append(response, book)
	}

	return &response, nil
}

func (mongoImpl *MongoGatewayImpl) GetSectionsByBookId(bookId string) (*[]models.BookSection, error)  {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	pipeline := make([]bson.D, 0, 0)
	queryPipeline := make([]bson.D, 0, 0)
	pipeline = append(pipeline, bson.D{{"$match", bson.D{{"_id", bookId}}}})
	pipeline = append(pipeline, bson.D{
		{"$lookup",
			bson.D{
				{"from", "bookSections"},
				{"localField", "content.sections.sectionId"},
				{"foreignField", "_id"},
				{"as", "index"},
			},
		},
	})
	pipeline = append(pipeline, bson.D{
		{"$project",
			bson.D{
				{"index", 1},
				{"_id", 0},
			},
		},
	})
	queryPipeline = append(queryPipeline, pipeline...)

	cursor, err := collection.Aggregate(ctx, queryPipeline)
	if err != nil {
		return nil, err
	}

	var response []models.BookSection
	var sections []models.BookSectionIndex
	err = cursor.All(ctx, &sections)
	if err != nil {
		return nil, err
	}

	for _, section := range sections[0].Index {
		response = append(response, section)
	}

	return &response, nil
}

func (mongoImpl *MongoGatewayImpl) GetBookSectionById(id string) (*models.BookSection, error) {
	var section *models.BookSection
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(bookSections)

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&section)
	if err != nil {
		return nil, err
	}

	return section, nil
}

func (mongoImpl *MongoGatewayImpl) GetBookById(id string) (*models.Book, error) {
	var book *models.Book
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return nil, errors.New("BOOK_NOT_FOUND")
	}

	return book, nil
}

func (mongoImpl *MongoGatewayImpl) GetBooksByAuthor(authorId string) (*[]models.Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	cursor, err := collection.Find(ctx, bson.M{"authors.authorId": authorId})
	if err != nil {
		return nil, err
	}

	var books []models.Book
	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, err
	}

	return &books, nil
}

func (mongoImpl *MongoGatewayImpl) GetBooksByCategory(category string) (*[]models.Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	collection := mongoImpl.client.Database(database).Collection(books)

	cursor, err := collection.Find(ctx, bson.D{
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
		return nil, err
	}

	var books []models.Book
	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, err
	}

	return &books, nil
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

func (mongoImpl *MongoGatewayImpl) UpdateBook(book *models.Book) (*models.Book, error) {
	var bookE *models.Book
	ctx, _ := context.WithTimeout(context.Background(), 30+time.Second)
	opts := options.Update().SetUpsert(true)
	collection := mongoImpl.client.Database(database).Collection(books)

	err := collection.FindOne(ctx, bson.M{"_id": book.Id}).Decode(&bookE)
	if err != nil {
		return nil, errors.New("BOOK_NOT_FOUND")
	}

	book.UpdatedAt = time.Now()

	_, err = collection.UpdateOne(ctx, bson.M{"_id": book.Id}, bson.D{{"$set", book}}, opts)
	if err != nil {
		return nil, err
	}

	return book, nil
}
