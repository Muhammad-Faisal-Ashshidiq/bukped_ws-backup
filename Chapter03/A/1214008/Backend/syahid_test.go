package syahid

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertKaryawan(t *testing.T) {
	dbname := "penggajian"
	karyawan := Karyawan{
		ID:      primitive.NewObjectID(),
		Nama:    "Ujang Saepudin",
		Status:  "Aktif",
		Jabatan: "Staff Administrasi",
		Gaji:    "RP 4.000.000",
	}
	insertedID := InsertKaryawan(dbname, karyawan)
	if insertedID == nil {
		t.Error("Failed to insert karyawan")
	}
}

func TestInsertHonor(t *testing.T) {
	dbname := "penggajian"
	honor := Honor{
		ID:      primitive.NewObjectID(),
		Nama:    "Asep Saepuloh",
		Status:  "Aktif",
		Jabatan: "Staff",
		Gaji:    "RP 1.500.000",
	}
	insertedID := InsertHonor(dbname, honor)
	if insertedID == nil {
		t.Error("Failed to insert honor")
	}
}

func TestInsertJob(t *testing.T) {
	dbname := "penggajian"
	job := Job{
		Namajob: "Staff Administrasi",
	}
	insertedID := InsertJob(dbname, job)
	if insertedID == nil {
		t.Error("Failed to insert job")
	}
}

func TestInsertTeam(t *testing.T) {
	dbname := "penggajian"
	team := Team{
		ID:        primitive.NewObjectID(),
		Nama:      "Uzumaki Memet",
		Deskripsi: "anjay mabar",
	}
	insertedID := InsertTeam(dbname, team)
	if insertedID == nil {
		t.Error("Failed to insert team")
	}
}

func TestInsertAbout(t *testing.T) {
	dbname := "penggajian"
	about := About{
		ID:       primitive.NewObjectID(),
		Isi_satu: "Ini isi satu",
		Isi_dua:  "Ini isi dua",
		Image:    "image.jpg",
	}
	insertedID := InsertAbout(dbname, about)
	if insertedID == nil {
		t.Error("Failed to insert about")
	}
}
