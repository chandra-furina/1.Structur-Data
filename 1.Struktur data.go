package main

import (
	"fmt"
	"time"
)

//1.Structur Data

type Mahasiswa struct {
	ID           int
	Nama         string
	TanggalLahir time.Time
	JurusanID    int
	Nilai        float64
	Status       string
}

type Jurusan struct {
	ID   int
	Nama string
}

type Aplikasipendaftaran struct {
	DaftarMahasiswa []Mahasiswa
	DaftarJurusan   []jurusan
}
