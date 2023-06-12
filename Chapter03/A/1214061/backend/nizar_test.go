package main

import (
	"testing"

	"github.com/example/corpo"
)

func TestCreateAndGetTagihanByID(t *testing.T) {
	// Simulate a new tagihan belanja data
	newTagihan := corpo.Tagihan{
		Nama:     "Tagihan Belanja",
		Nominal:  1000000,
		Tanggal:  "2022-03-05",
		Dibayar:  false,
		Pemilik:  "PT. XYZ",
		Kategori: "Belanja",
	}

	// Create new tagihan belanja
	id := corpo.CreateTagihan(newTagihan)

	// Get tagihan belanja by ID
	result, err := corpo.GetTagihanByID(id.(string))
	if err != nil {
		t.Errorf("Error while getting tagihan belanja: %v", err)
	}

	// Check if result is equal to the newTagihan data
	if result.Nama != newTagihan.Nama {
		t.Errorf("Expected nama to be %s, but got %s", newTagihan.Nama, result.Nama)
	}
	if result.Nominal != newTagihan.Nominal {
		t.Errorf("Expected nominal to be %d, but got %d", newTagihan.Nominal, result.Nominal)
	}
	if result.Tanggal != newTagihan.Tanggal {
		t.Errorf("Expected tanggal to be %s, but got %s", newTagihan.Tanggal, result.Tanggal)
	}
	if result.Dibayar != newTagihan.Dibayar {
		t.Errorf("Expected dibayar to be %t, but got %t", newTagihan.Dibayar, result.Dibayar)
	}
	if result.Pemilik != newTagihan.Pemilik {
		t.Errorf("Expected pemilik to be %s, but got %s", newTagihan.Pemilik, result.Pemilik)
	}
	if result.Kategori != newTagihan.Kategori {
		t.Errorf("Expected kategori to be %s, but got %s", newTagihan.Kategori, result.Kategori)
	}
}
