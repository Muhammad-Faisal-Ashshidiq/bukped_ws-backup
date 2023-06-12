package rizki

import (
	"fmt"
	"testing"
)

func TestInsertMahasiswa(t *testing.T) {
	nm := "Farhan Rizki Maulana"
	npm := "1214020"
	smt := "4"
	kls := "2A"
	prd := "D4 Teknik Informatika"
	hasil := InsertMahasiswa(nm, npm, smt, kls, prd)
	fmt.Println(hasil)

}

func TestInsertPresensi(t *testing.T) {
	khd := "Hadir"
	ktr := "Masuk"
	hasil := InsertPresensi(khd, ktr)
	fmt.Println(hasil)
}

func TestInsertMataKuliah(t *testing.T) {
	kdmk := "TI41264"
	nmmk := "Pemrograman 3"
	sks := "3"
	jrs := "Teknik Informatika"
	hasil := InsertMataKuliah(kdmk, nmmk, sks, jrs)
	fmt.Println(hasil)
}

func TestInsertJadwalKuliah(t *testing.T) {
	matkul := "Pemrograman 3"
	hari := "Senin"
	jmulai := "13:00"
	jselesai := "18:00"
	ruangan := "Lab 314"
	hasil := InsertJadwalKuliah(matkul, hari, jmulai, jselesai, ruangan)
	fmt.Println(hasil)
}

func TestInsertDosen(t *testing.T) {
	nidn := "0410118609"
	nama := "Rolly Maulana Awangga,S.T.,MT.,CAIP, SFPC."
	matkul := "Pemrograman 3"
	hasil := InsertDosen(nidn, nama, matkul)
	fmt.Println(hasil)
}
