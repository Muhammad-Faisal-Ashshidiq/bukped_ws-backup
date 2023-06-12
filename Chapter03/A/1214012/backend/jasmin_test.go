package jasmin_test

import (
	"fmt"
	"testing"
)

func TestInsertNilai(t *testing.T) {
	ID := 1
	nama_matkul := Network Programming
	kode_matkul := "TI42253"
	sks := "3"
	grade := "A"
	matakuliah := dhs{
		Nama_matkul: "Network Programming",
		Kode_matkul: "TI42253",
		Nama_dosen:  "M. Yusril Helmi",
	}
	hasil := InsertNilai(ID, nama_matkul, kode_matkul, sks, grade, matakuliah)
	fmt.Println(hasil)

}

func TestGetdhsFromKode_matkul(t *testing.T) {
	kode_matkul := "TI42253"
	matakuliah := GetDhsFromKode_matkul(kode_matkul)
	fmt.Println(nilai)
}
