package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Mongostring = "mongodb+srv://GoProduct:goMux123@cluster0.xdoq5.mongodb.net/?retryWrites=true&w=majority"

var Collection1 *mongo.Collection
var Collection2 *mongo.Collection

const dbName = "goProducts"
const colName1 = "users"
const colName2 = "products"

func Connect() {
	fmt.Println("Let us create a connection with mongodb atlas")
	clientOption := options.Client().ApplyURI(Mongostring)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established")

	Collection1 = client.Database(dbName).Collection(colName1)
	Collection2 = client.Database(dbName).Collection(colName2)

	fmt.Println("Users instance is created")
	fmt.Println("Products instance is also created")

}
