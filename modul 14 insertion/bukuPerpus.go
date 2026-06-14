package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku struct {
	Pustaka  [nMax]Buku
	nPustaka int
}

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka.Pustaka[i].id)
		fmt.Scan(&pustaka.Pustaka[i].judul)
		fmt.Scan(&pustaka.Pustaka[i].penulis)
		fmt.Scan(&pustaka.Pustaka[i].penerbit)
		fmt.Scan(&pustaka.Pustaka[i].eksemplar)
		fmt.Scan(&pustaka.Pustaka[i].tahun)
		fmt.Scan(&pustaka.Pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	max := 0
	for i := 0; i < n; i++ {
		if pustaka.Pustaka[i].rating > pustaka.Pustaka[max].rating {
			max = i
		}
	}
	fmt.Println("=== Buku Terfavorit ===")
	fmt.Println("Judul   :", pustaka.Pustaka[max].judul)
	fmt.Println("Penulis :", pustaka.Pustaka[max].penulis)
	fmt.Println("Penerbit:", pustaka.Pustaka[max].penerbit)
	fmt.Println("Tahun   :", pustaka.Pustaka[max].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		j := i
		temp := pustaka.Pustaka[j]
		for j > 0 && pustaka.Pustaka[j-1].rating < temp.rating {
			pustaka.Pustaka[j] = pustaka.Pustaka[j-1]
			j--
		}
		pustaka.Pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("=== Buku Rating Tertinggi ===")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Println(i+1,".", pustaka.Pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 0
	kanan := n - 1
	ketemu := -1
	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		if pustaka.Pustaka[mid].rating == r {
			ketemu = mid
			break
		} else if pustaka.Pustaka[mid].rating < r {
			kanan = mid - 1
		} else {
			kiri = mid + 1
		}
	}
	fmt.Println("=== Hasil Pencarian Rating", r, "===")
	if ketemu == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		fmt.Println("Judul    :", pustaka.Pustaka[ketemu].judul)
		fmt.Println("Penulis  :", pustaka.Pustaka[ketemu].penulis)
		fmt.Println("Penerbit :", pustaka.Pustaka[ketemu].penerbit)
		fmt.Println("Tahun    :", pustaka.Pustaka[ketemu].tahun)
		fmt.Println("Eksemplar:", pustaka.Pustaka[ketemu].eksemplar)
	}
}

func main() {
	var daftar DaftarBuku
	var N int
	var r int

	fmt.Scan(&N)
	DaftarkanBuku(&daftar, N)
	CetakTerfavorit(daftar, N)
	UrutBuku(&daftar, N)
	Cetak5Terbaru(daftar, N)
	fmt.Println()
	fmt.Print("Masukkan rating yang ingin dicari: ")
	fmt.Scan(&r)
	fmt.Println()
	CariBuku(daftar, N, r)
}