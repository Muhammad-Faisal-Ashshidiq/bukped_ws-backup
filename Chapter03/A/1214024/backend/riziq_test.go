package riziq

import (
	"fmt"
	"testing"
)

func TestInsertPelanggan(t *testing.T) {
	nm := "Farhan Riziq"
	alm := "bandung"
	tlp := "081238765694"
	em := "farhan@gmail.com"
	hasil := InsertPelanggan(nm, alm, tlp, em)
	fmt.Println(hasil)

}

func TestInsertTagihan(t *testing.T) {
	tgh := "08.01.2021"
	tth := "$5000.00"
	sp := "complete"
	hasil := InsertTagihan(tgh, tth, sp)
	fmt.Println(hasil)
}

func TestInsertPembayaran(t *testing.T) {
	tby := "05.01.2021"
	jby := "$2500.00"
	mby := "transfer"
	hasil := InsertPembayaran(tby, jby, mby)
	fmt.Println(hasil)
}

func TestInsertProduk(t *testing.T) {
	npr := "Spotify"
	hpr := "$5000.00"
	hasil := InsertProduk(npr, hpr)
	fmt.Println(hasil)
}

func TestInsertAbout(t *testing.T) {
	satu := "isi satu"
	dua := "isi dua"
	hasil := InsertAbout(satu, dua)
	fmt.Println(hasil)
}
