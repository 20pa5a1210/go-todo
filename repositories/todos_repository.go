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

func (todo *TodoRepository) AddTodo(userId string, todos models.Todos) (models.Todos, error) {
	newTodo := models.Todos{
		ID:   primitive.NewObjectID(),
		Text: todos.Text,
	}
	filter := bson.M{"email": userId}
	update := bson.M{"$push": bson.M{"todos": newTodo}}

	_, err := todo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return models.Todos{}, err
	}
	updateUser := models.Todos{}
	err = todo.collection.FindOne(context.Background(), filter).Decode(&updateUser)
	if err != nil {
		return models.Todos{}, err
	}

	return updateUser, nil
}
