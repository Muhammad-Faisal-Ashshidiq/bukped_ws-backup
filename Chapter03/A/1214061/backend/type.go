package main

import (
	"time"
)

type Tagihan struct {
	ID           int
	TotalHarga   float64
	TanggalBeli  time.Time
	Barang       []Barang
	Pembeli      Pembeli
	StatusBayar  bool
	TanggalBayar time.Time
}

type Barang struct {
	NamaBarang string
	Harga      float64
}

type Pembeli struct {
	Nama     string
	Email    string
	NoTelp   string
	Alamat   string
	KodePos  string
	Kota     string
	Provinsi string
}
