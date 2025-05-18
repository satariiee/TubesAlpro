package main

import (
	"fmt"
)

const NMAX int = 1000

type sampah struct {
	jenis  string
	jumlah int
}

type dataSampah [NMAX]sampah

func main() {
	var data dataSampah
	fmt.Println("\n--- Aplikasi Pengelolaan Sampah ---")
	fmt.Println("1. Tambah Data Sampah")
	fmt.Println("2. Ubah Data Sampah")
	fmt.Println("3. Hapus Data Sampah")
	fmt.Println("4. Cari Data Sampah")
	fmt.Println("5. Urutkan Data Sampah")
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
	var i, idx int
	var found bool = false
	for i = 0; i < NMAX && !found; i++ {
		if A[i].jenis == "" {
			index = i
			found = true
		}
	}

	if found {
		fmt.Print("Masukan jenis sampah: ")
		fmt.Scan(&A[idx].jenis)
		fmt.Print("Masukan jumlah sampah: ")
		fmt.Scan(&A[idx].jumlah)
		fmt.Print("Masukan berat sampah: ")
		fmt.Scan(&A[idx].berat)
		fmt.Println("Data sampah berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas penyimpanan data sampah penuh.")
	}

}
