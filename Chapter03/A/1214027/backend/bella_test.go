package bella

import (
	"fmt"
	"testing"
)

func TestInsertMataKuliah(t *testing.T) {
	namamatakuliah := "Pemrograman III"
	kodematakuliah := "TI41264"
	dosen := "ROLLY MAULANA AWANGGA"
	sks := "3"
	hasil := InsertMataKuliah(namamatakuliah, kodematakuliah, dosen, sks)
	fmt.Println(hasil)
}

func TestInsertJadwalKuliah(t *testing.T) {
	hari := "Senin"
	jammulai := "13:00"
	jamselesai := "18:00"
	ruang := "201"
	hasil := InsertJadwalKuliah(hari, jammulai, jamselesai, ruang)
	fmt.Println(hasil)
}

func TestInsertKelas(t *testing.T) {
	ruang := "201"
	kapasitasmhs := "40"
	hasil := InsertKelas(ruang, kapasitasmhs)
	fmt.Println(hasil)
}

func TestInsertDosen(t *testing.T) {
	namadosen := "ROLLY MAULANA AWANGGA"
	kodedosen := "TI41264"
	matakuliah := "Pemrograman III"
	hasil := InsertDosen(namadosen, kodedosen, matakuliah)
	fmt.Println(hasil)
}

func TestInsertMahasiswa(t *testing.T) {
	namamhs := "GABRIELLA YOUZANNA RORONG"
	kelas := "2A"
	prodi := "D4 Teknik Informatika"
	hasil := InsertMahasiswa(namamhs, kelas, prodi)
	fmt.Println(hasil)
}