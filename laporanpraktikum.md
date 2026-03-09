# <h1>Laporan Praktikum Modul 1 - 2</h1>
<p align-"center">[Arsal Aji Ngroho] - [109082530039]</p>

### 1. [SOAL 1]
#### Soal1.go

    package main

		import "fmt"

		func main() {

	var (
	satu, dua, tiga string
		temp string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
		fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
			temp = satu
			satu = dua
			dua = tiga
			tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}


### Output Unguided : 

##### Output
![Screenshot Output Unguided 1_1](gambar/soal-1.PNG)
<br>[Program meminta untuk memasukkan tiga string dengan Output awal satu, dua dan tiga kemudian masuk pada operasinya dimana ada temp untuk menyimpan variabel sementara(variabel satu), ini berfungsi agar saat print output variabel satu tidak hilang. setelah masuk ke operasinya, angka satu akan langsung berubah menjadi dua, kemudia angka dua menjadi tiga, lalu tiga menjadi satu. Dan hasil Output Akhirnya adalah dua, tiga, satu]

### 2. [SOAL 2]
#### Soal2.go

    package main

    import "fmt"

        func main() {
	    var a, b, c, d string
	
	hasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("percobaan ", i, ": ")
		fmt.Scan(&a, &b, &c, &d)
		if a != "merah" || b != "kuning" || c != "hijau" || d != "ungu" {
			hasil = false
		}
	}
	fmt.Println(hasil)
}

