// db/db.go
package database

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

var (
    client     *mongo.Client
    dbName     = "yourdbname"
    collection = "yourcollection"
)

// Connect initializes a connection to the MongoDB database
func Connect() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://awaismumtaz0099:25213291231919@cluster0.3so1bcq.mongodb.net")

    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")

    // You can use the "client" variable to perform MongoDB operations.
    // For example: client.Database(dbName).Collection(collection).Find(...)

}

// GetClient returns the MongoDB client
func GetClient() *mongo.Client {
    return client
}
