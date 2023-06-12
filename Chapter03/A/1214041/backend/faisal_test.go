package faisal

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "Profil",
}

var MongoConn = atdb.MongoConnect(MongoInfo)


// func TestInsertProfile(t *testing.T) {
// 	pendidikan := "S1"
// 	username := "Fluxy"
// 	checkin := "masuk"
// 	biodata := Profile{
// 		FullName : "Fitrana Soleh",
// 		Email : "fitrana12@gmail.com",
// 	}
// 	hasil:=InsertOneDoc(username , checkin , biodata)
// 	fmt.Println(hasil)
// 	fmt.Println(pendidikan)
	

// }

// func TestGetProfileByUsername(t *testing.T) {
// 	username := "Fitrana "
// 	biodata:=GetProfileByUsername(username)
// 	fmt.Println(biodata)
// }

func TestInsertProfile(t *testing.T) {
	username := "Fitrana Soleh"
	bio := "sibuk"
	pendidikan := "S1"
	result := InsertProfile(pendidikan, username, bio, "masuk", Profile{FullName: "Fitrana Soleh", Email: "fitrana12@gmail.com"}, MongoConn)
	fmt.Println(result)
}
