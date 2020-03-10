package main

import "context"
import "os"
import "fmt"
import "time"
import "go.mongodb.org/mongo-driver/bson"
import _"go.mongodb.org/mongo-driver/bson/primitive"
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"
import "go.mongodb.org/mongo-driver/mongo/readpref"

func firstone(){
    // 最初の1件

    // 疑問。Contextとは何をするものでしょうか？
    //ctx, _ := context.Background()
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Clinet作成
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
    if err != nil {
		fmt.Println("ConnectError", err)
		os.Exit(1)
	}
	// Ping?
	err = client.Ping(ctx, readpref.Primary())
    if err != nil {
		fmt.Println("PingError", err)
		os.Exit(1)
	}
	// 終了時の処理定義
	defer client.Disconnect(ctx)

	var doc bson.Raw
	// Collectionの作成
    col := client.Database("mydb").Collection("CAR")
	filter := bson.D{}
	opt := options.FindOne()
	err = col.FindOne(context.Background(),filter, opt).Decode(&doc)
    if err != nil {
		fmt.Println("FindOneError", err)
		os.Exit(1)
	}
	fmt.Println("CAR First One:", doc.String())
}

func findall( ){
    // all件
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
    if err != nil {
		fmt.Println("ConnectError", err)
		os.Exit(1)
	}
    col := client.Database("mydb").Collection("CAR")
	// 検索結果のカーソル
	cur, err := col.Find(context.Background(),bson.D{})
    if err != nil {
		fmt.Println("FindError", err)
		os.Exit(1)
	}
	// 終了時処理でカーソルクローズ
	defer cur.Close(ctx)

	for cur.Next(ctx){
		var result bson.Raw
		err := cur.Decode(&result)
		if err != nil {
		    fmt.Println("DecodeError", err)
	    }
	    fmt.Println("CAR All:", result.String())
	}
}

func main(){
	firstone()
	findall()
}
