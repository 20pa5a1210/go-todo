package repositories

import (
	"context"
	"log"

	"github.com/20pa5a1210/go-todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.jhj7ej8.mongodb.net/test")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("test")
	collection := db.Collection("users")
	log.Print("DB Connected")

	return &UserRepository{
		collection: collection,
	}
}

func (ur *UserRepository) CreateUser(user models.User) (models.User, error) {
	result, err := ur.collection.InsertOne(context.Background(), user)
	if err != nil {
		return models.User{}, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	filter := bson.M{"email": email}
	err := ur.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// doc not found
			return models.User{}, nil
		}
		return models.User{}, nil
	}
	return user, nil
}

func (ur *UserRepository) GetUserByID(userId string) (models.User, error) {
	var user models.User
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return models.User{}, err
	}
	filter := bson.M{"_id": objId}
	err = ur.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) GetUsers() ([]models.User, error) {
	cursor, err := ur.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
