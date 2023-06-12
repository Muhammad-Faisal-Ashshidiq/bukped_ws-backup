package gipar

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertMahasiswa(t *testing.T) {
	mhs := Mahasiswa{
		ID:          primitive.NewObjectID(),
		NamaLengkap: "John Doe",
		NPM:         "12345",
		Alamat:      "Jalan Raya 123",
		NilaiTugas:  80.0,
		NilaiUTS:    85.0,
		NilaiUAS:    90.0,
		Grade:       "A",
	}
	insertedID := InsertMahasiswa("nilai", mhs)
	if insertedID == nil {
		t.Errorf("Error inserting data into Mahasiswa collection")
	}
}

func TestInsertSyaratGrade(t *testing.T) {
	sg := SyaratGrade{
		ID:     primitive.NewObjectID(),
		Nilai:  80.0,
		Syarat: ">= 80",
		Grade:  "A",
	}
	insertedID := InsertSyaratGrade("nilai", sg)
	if insertedID == nil {
		t.Errorf("Error inserting data into SyaratGrade collection")
	}
}

func TestInsertPrestasi(t *testing.T) {
	ps := Prestasi{
		ID:          primitive.NewObjectID(),
		NamaLengkap: "John Doe",
		NPM:         "12345",
		Prestasi:    "Juara 1 Lomba Cerdas Cermat",
	}
	insertedID := InsertPrestasi("nilai", ps)
	if insertedID == nil {
		t.Errorf("Error inserting data into Prestasi collection")
	}
}

func TestInsertNilai(t *testing.T) {
	n := Nilai{
		ID:          primitive.NewObjectID(),
		MahasiswaID: primitive.NewObjectID(),
		NPM:         "12345",
		AlamatID:    primitive.NewObjectID(),
		NilaiTugas:  80.0,
		NilaiUTS:    85.0,
		NilaiUAS:    90.0,
	}
	insertedID := InsertNilai("nilai", n)
	if insertedID == nil {
		t.Errorf("Error inserting data into Nilai collection")
	}
}

func TestInsertAlamat(t *testing.T) {
	alm := Alamat{
		ID:          primitive.NewObjectID(),
		MahasiswaID: primitive.NewObjectID(),
		NamaLengkap: "John Doe",
		Alamat:      "Jalan Raya 123",
	}
	insertedID := InsertAlamat("nilai", alm)
	if insertedID == nil {
		t.Errorf("Error inserting data into Alamat collection")
	}
}
