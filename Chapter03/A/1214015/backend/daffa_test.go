package daffa

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertUser(t *testing.T) {
	dbname := "proyek-2"
	user := User{
		ID:     primitive.NewObjectID(),
		Nama:   "Arya",
		Gender: "Perempuan",
		Email:  "Arya@gmail.com",
	}
	insertedID := InsertUser(dbname, user)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestInsertPendaftaran(t *testing.T) {
	dbname := "proyek-2"
	pendaftaran := Pendaftaran{
		ID:         primitive.NewObjectID(),
		Nama_siswa: "Arya",
		Nis:        "92348927348",
		Nik:        "72619873972094293",
	}
	insertedID := InsertPendaftaran(dbname, pendaftaran)
	if insertedID == nil {
		t.Error("Failed to insert pendaftaran")
	}
}

func TestInsertPembayaran(t *testing.T) {
	dbname := "proyek-2"
	pembayaran := Pembayaran{
		ID:          primitive.NewObjectID(),
		Status:      "Dibayar",
		Total_bayar: "200,000",
	}
	insertedID := InsertPembayaran(dbname, pembayaran)
	if insertedID == nil {
		t.Error("Failed to insert pembayaran")
	}
}
func TestInsertPengumuman(t *testing.T) {
	dbname := "proyek-2"
	pengumuman := Pengumuman{
		ID:            primitive.NewObjectID(),
		Hasil_seleksi: "LULUS",
		Nilai:         "97",
		Program:       "Fighting",
	}
	insertedID := InsertPengumuman(dbname, pengumuman)
	if insertedID == nil {
		t.Error("Failed to insert pengumuman")
	}
}
func TestInsertKursus(t *testing.T) {
	dbname := "proyek-2"
	kursus := Kursus{
		ID:             primitive.NewObjectID(),
		Nama_kursus:    "FIGHTING GAME",
		Jenjang_kursus: "Program 3",
		Pengajar:       "Jon Snow",
	}
	insertedID := InsertKursus(dbname, kursus)
	if insertedID == nil {
		t.Error("Failed to insert kursus")
	}
}
