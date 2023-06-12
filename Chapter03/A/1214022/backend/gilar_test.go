package ckbackend

import (
	"fmt"
	"testing"
)

// func TestInsertDashboard(t *testing.T) {
// 	dbname := "tablerps"
// 	dashboard := Dashboard{
// 		ID:         primitive.NewObjectID(),
// 		Username:   "Patricia Semklo",
// 		Email:      "patricia.semklo@app.com",
// 		Location:   "Denmark",
// 		Orders:     "6",
// 		Lastorders: "#904300",
// 		Totalspent: "$3,455.00",
// 		Sks:        "12",
// 	}
// 	insertedID := InsertDashboard(dbname, dashboard)
// 	if insertedID == nil {
// 		t.Error("Failed to insert dashboard")
// 	}
// }

// func TestInsertMahasiswa(t *testing.T) {
// 	dbname := "tablerps"
// 	mahasiswa := Mahasiswa{
// 		Nama_mhs:   "John Doe",
// 		Email_mhs:  "johndoe@flex.co",
// 		Gambar_mhs: "image.jpg",
// 	}
// 	insertedID := InsertMahasiswa(dbname, mahasiswa)
// 	if insertedID == nil {
// 		t.Error("Failed to insert mahasiswa")
// 	}
// }

// func TestInsertDosen(t *testing.T) {
// 	dbname := "tablerps"
// 	dosen := Dosen{
// 		Nama:      "Macauley Herring",
// 		Jabatan:   "DOSEN",
// 		Noted:     "Dance is the hidden language of the soul.",
// 		Img_dosen: "image.jpg",
// 	}
// 	insertedID := InsertDosen(dbname, dosen)
// 	if insertedID == nil {
// 		t.Error("Failed to insert dosen")
// 	}
// }

//	func TestInsertAbout(t *testing.T) {
//		dbname := "tablerps"
//		about := About{
//			Pertanyaan: "Apakah kurikulum ini menyulitkan hidup?",
//			Jawaban:    "jelas ini akan membuat jiwamu bergetar sampai ke bulan",
//		}
//		insertedID := InsertAbout(dbname, about)
//		if insertedID == nil {
//			t.Error("Failed to insert about")
//		}
//	}
//
//	func TestInsertContacus(t *testing.T) {
//		dbname := "tablerps"
//		contacus := Contacus{
//			Phone_number: "0822126722",
//			Email:        "kucingterbang@gmail.crot",
//		}
//		insertedID := InsertContacus(dbname, contacus)
//		if insertedID == nil {
//			t.Error("Failed to insert team")
//		}
//	}

func TestGetDataDosen(t *testing.T) {
	stats := "DOSEN"
	data := GetDataDosen(stats)
	fmt.Println(data)
}

func TestGetDataAbout(t *testing.T) {
	stats := "Apakah kurikulum ini menyulitkan hidup?"
	data := GetDataAbout(stats)
	fmt.Println(data)
}

func TestGetDataContacus(t *testing.T) {
	stats := "0822126722"
	data := GetDataContacus(stats)
	fmt.Println(data)
}

func TestGetDataDashboard(t *testing.T) {
	stats := "Denmark"
	data := GetDataDashboard(stats)
	fmt.Println(data)
}
func TestGetDataMahasiswa(t *testing.T) {
	stats := "johndoe@flex.co"
	data := GetDataMahasiswa(stats)
	fmt.Println(data)
}
