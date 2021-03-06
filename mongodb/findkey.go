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

func findkey( ){
    // 条件検索
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx,
	               options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
    if err != nil {
		fmt.Println("ConnectError", err)
		os.Exit(1)
	}
	// コレクション
    col := client.Database("mydb").Collection("CAR")
	// 検索条件、結果複数
	filter := bson.M{"name": "subaru"}
	// 検索結果のカーソル
	cur, err := col.Find(context.Background(), filter)
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
	    fmt.Println("CAR Find by Key:", result.String())
	}
}

func main(){
	findkey()
}
