package corpo

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) interface{} {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.Background(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
		return nil
	}
	return insertResult.InsertedID
}

func InsertTagihan(pembeli string, daftarBarang []Barang) interface{} {
	var tagihan Tagihan
	tagihan.Pembeli = pembeli
	tagihan.DaftarBarang = daftarBarang
	tagihan.TanggalTagihan = primitive.NewDateTimeFromTime(time.Now().UTC())
	return InsertOneDoc("mydatabase", "tagihan", tagihan)
}

func GetTagihanByPembeli(pembeli string) []Tagihan {
	var tagihans []Tagihan
	collection := MongoConnect("mydatabase").Collection("tagihan")
	filter := bson.M{"pembeli": pembeli}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Printf("GetTagihanByPembeli: %v\n", err)
		return nil
	}
	for cur.Next(context.Background()) {
		var tagihan Tagihan
		err := cur.Decode(&tagihan)
		if err != nil {
			fmt.Printf("GetTagihanByPembeli: %v\n", err)
			return nil
		}
		tagihans = append(tagihans, tagihan)
	}
	if len(tagihans) == 0 {
		return nil
	}
	return tagihans
}
