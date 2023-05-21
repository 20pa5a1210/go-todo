package repositories

import (
	"context"
	"errors"
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

func (todo *TodoRepository) CreateTodoInstance(instance models.Todos) (models.Todos, error) {
	result, err := todo.collection.InsertOne(context.Background(), instance)
	if err != nil {
		return models.Todos{}, err
	}
	instance.ID = result.InsertedID.(primitive.ObjectID)
	return instance, nil
}

func (todo *TodoRepository) AddTodo(userId string, todos models.Todo) (models.Todos, error) {
	newTodo := models.Todo{
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

func (todo *TodoRepository) GetTodos(userId string) ([]models.Todo, error) {
	var user models.Todos
	filter := bson.M{"email": userId}
	err := todo.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []models.Todo{}, err
		}
		return nil, err
	}
	return user.Todos, nil
}

func (todo *TodoRepository) DeleteTodo(todoId string, email string) error {
	filter := bson.M{"email": email}
	update := bson.M{
		"$pull": bson.M{
			"todos": bson.M{
				"_id": todoId,
			},
		},
	}
	updated, err := todo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if updated.ModifiedCount == 0 {
		return errors.New("todo not found")
	}
	return nil
}
