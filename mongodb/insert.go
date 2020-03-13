package main

import "context"
import "os"
import "fmt"
import "time"
import "go.mongodb.org/mongo-driver/bson"
import _"go.mongodb.org/mongo-driver/bson/primitive"
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"
import _"go.mongodb.org/mongo-driver/mongo/readpref"

func insertBsonD(col *mongo.Collection) error {
	// 順番どおりに入る
	bsonD := bson.D{ {"name", "ford"},
	                 {"since", 1903},
					 {"updateat", time.Now()}}
    fmt.Println("carType bsonD", bsonD)
	_, err := col.InsertOne(context.Background(), bsonD)
	return err
}

func insertBsonM(col *mongo.Collection) error {
	// 順不同で入る
	bsonM := bson.M{ "name": "Daimler",
	                 "since": 1926,
					 "updateat": time.Now()}
    fmt.Println("carType bsonM", bsonM)
	_, err := col.InsertOne(context.Background(), bsonM)
	return err
}

type carType struct {
	nameSt string
	name string
	sinceSt string
	since int
	updataatSt string
	udateat time.Time
}

func insertStruct(col *mongo.Collection) error {
	// 構造体が入らない誰か教えて。
	doc := carType{ "name", "Bmw", "since", 1919, "udateat", time.Now(), }
    fmt.Println("carType", doc)
	_, err := col.InsertOne(context.Background(), doc)
	return err
}

func main() {
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
    if err != nil {
        fmt.Println("ConnectError", err)
        os.Exit(1)
    }
    defer client.Disconnect(context.Background())
	col := client.Database("mydb").Collection("CAR")

	err = insertBsonD(col)
    if err != nil {
        fmt.Println("insertBsonD()Error", err)
        os.Exit(1)
    }
	err = insertBsonM(col)
    if err != nil {
        fmt.Println("insertBsonM()Error", err)
        os.Exit(1)
    }
	err = insertStruct(col)
    if err != nil {
        fmt.Println("insertStruct()Error", err)
        os.Exit(1)
    }
}



