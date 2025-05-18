package main

import (
	"fmt"
)

const NMAX int = 1000

type sampah struct {
	jenis           string
	jumlah          int
	daurUlang       bool
	metodeDaurUlang string
}

type dataSampah [NMAX]sampah

func main() {
	var data dataSampah
	fmt.Println("\n--- Aplikasi Pengelolaan Sampah ---")
	fmt.Println("1. Tambah Data Sampah")
	fmt.Println("2. Ubah Data Sampah")
	fmt.Println("3. Hapus Data Sampah")
	fmt.Println("4. Cari Sampah")
	fmt.Println("5. Urutkan Sampah")
	fmt.Println("6. Tampilkan Statistik")
	fmt.Println("0. Keluar")

	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tambahData(&data)
	}
}

func tambahData(A *dataSampah) {
	fmt.Scan(A)
}
