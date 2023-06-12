package syahid

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertKaryawan(db string, karyawan Karyawan) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("karyawan").InsertOne(context.TODO(), karyawan)
	if err != nil {
		fmt.Printf("InsertKaryawan: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertHonor(db string, honor Honor) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("honor").InsertOne(context.TODO(), honor)
	if err != nil {
		fmt.Printf("InsertHonor: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertJob(db string, job Job) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("job").InsertOne(context.TODO(), job)
	if err != nil {
		fmt.Printf("InsertJob: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertTeam(db string, team Team) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("team").InsertOne(context.TODO(), team)
	if err != nil {
		fmt.Printf("InsertTeam: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertAbout(db string, about About) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("about").InsertOne(context.TODO(), about)
	if err != nil {
		fmt.Printf("InsertAbout: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetKaryawan(db string, id interface{}) (*Karyawan, error) {
	var result Karyawan
	err := MongoConnect(db).Collection("karyawan").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetHonor(db string, id interface{}) (*Honor, error) {
	var result Honor
	err := MongoConnect(db).Collection("honor").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetJob(db string, id interface{}) (*Job, error) {
	var result Job
	err := MongoConnect(db).Collection("job").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetTeam(db string, id interface{}) (*Team, error) {
	var result Team
	err := MongoConnect(db).Collection("team").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetAbout(db string, id interface{}) (*About, error) {
	var result About
	err := MongoConnect(db).Collection("about").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
