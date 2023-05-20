package repositories

import (
	"context"
	"log"

	"github.com/20pa5a1210/go-todo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository() *TodoRepository {
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.jhj7ej8.mongodb.net/test")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("test")
	collection := db.Collection("todos")
	log.Print("DB Connected")

	return &TodoRepository{
		collection: collection,
	}
}

func (todo *TodoRepository) CreateTodoInstance(instance models.Todo) (models.Todo, error) {
	result, err := todo.collection.InsertOne(context.Background(), instance)
	if err != nil {
		return models.Todo{}, err
	}
	instance.Id = result.InsertedID.(primitive.ObjectID)
	return instance, nil
}
