package todoDAO

import (
	"GOLANG/todo/db"
	"GOLANG/todo/forms"
	guidmdl "GOLANG/todo/utility"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type TaskModel struct{}

var server = "127.0.0.1"

var client = db.GetClient()

func (m *TaskModel) Create(data forms.CreateTaskCommand) error {

	var client = db.GetClient()
	var ctx = context.Background()
	collection := client.Database("todo").Collection("tasks")
	data.Id = guidmdl.GetGUID()
	_, err := collection.InsertOne(context.TODO(), bson.M{"id": data.Id, "title": data.Title, "description": data.Description, "completed": data.Completed})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	return err
}

func (m *TaskModel) Get() (list []*forms.CreateTaskCommand, err error) {

	var client = db.GetClient()
	var ctx = context.Background()
	collection := client.Database("todo").Collection("tasks")
	cur, err := collection.Find(context.TODO(), bson.M{"completed": false})
	for cur.Next(context.TODO()) {
		var li forms.CreateTaskCommand
		err = cur.Decode(&li)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		list = append(list, &li)
	}
	defer client.Disconnect(ctx)

	return list, err
}

func (m *TaskModel) Update(id string, data forms.CreateTaskCommand) (err error) {
	var client = db.GetClient()
	var ctx = context.Background()
	collection := client.Database("todo").Collection("tasks")
	filter := bson.M{"id": id}

	update := bson.M{"$set": bson.M{"title": data.Title, "description": data.Description, "completed": data.Completed}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return err
}

func (m *TaskModel) Delete(id string) (err error) {

	var client = db.GetClient()
	var ctx = context.Background()
	collection := client.Database("todo").Collection("tasks")
	filter := bson.M{"id": id}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	return err
}

func (m *TaskModel) GetTodoByIdDAO(id string) (forms.CreateTaskCommand, error) {
	var client = db.GetClient()
	var ctx = context.Background()
	collection := client.Database("todo").Collection("tasks")
	filter := bson.M{"id": id}
	var todo forms.CreateTaskCommand
	fetchError := collection.FindOne(context.TODO(), filter).Decode(&todo)
	if fetchError != nil {
		return todo, fetchError
	}
	defer client.Disconnect(ctx)

	return todo, nil
}
