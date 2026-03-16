# <h1 align="center"> Laporan Praktikum Modul 3 </h1>
<p align="center">  [Arsal Aji Nugroho] - [109082530039] </p>

## Unguided 

### 1. [PERMUTASI_DAN_KOMBINASI]
#### Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. 
#### Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. 
#### Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
#### Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!
#### P(n, r) = n! / (n−r)!, sedangkan C(n, r) = n! / r!(n−r)!

```go
    package main

    import "fmt"

    func faktorial(n int) int {
	    hasil := 1
	    for i := 1; i <= n; i++ {
		hasil *= i
	    }
	    return hasil
    }

    func permutasi(n, r int) int {
	    return faktorial(n) / faktorial(n-r)
    }

    func kombinasi(n, r int) int {
	    return faktorial(n) / (faktorial(r) * faktorial(n-r))
    }

    func main() {
	    var a, b, c, d int
	        fmt.Print("Masukan nilai n: ")
	        fmt.Scanln(&a)
	        fmt.Print("Masukan nilai r: ")
	        fmt.Scanln(&b)
	        fmt.Print("Masukan nilai n untuk kombinasi: ")
	        fmt.Scanln(&c)
	        fmt.Print("Masukan nilai r untuk kombinasi: ")
	        fmt.Scanln(&d)

	        fmt.Println("Permutasi: ", permutasi(a, c),"   ", "Kombinasi: ", kombinasi(a, c))
	        fmt.Println("Permutasi: ", permutasi(b, d), "    ", "Kombinasi: ", kombinasi(b, d))
    }

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](gambar/permutasi.JPG)
[]

### 2. [FoG]
#### Diberikan tiga buah fungsi matematika yaitu f (x) = x^2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function.
#### Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.
#### Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!

```go
    package main

    import "fmt"

    func faktorial(n int) int {
	    hasil := 1
	    for i := 1; i <= n; i++ {
		hasil *= i
	    }
	    return hasil
    }

    func permutasi(n, r int) int {
	    return faktorial(n) / faktorial(n-r)
    }

    func kombinasi(n, r int) int {
	    return faktorial(n) / (faktorial(r) * faktorial(n-r))
    }

    func main() {
	    var a, b, c, d int
	        fmt.Print("Masukan nilai n: ")
	        fmt.Scanln(&a)
	        fmt.Print("Masukan nilai r: ")
	        fmt.Scanln(&b)
	        fmt.Print("Masukan nilai n untuk kombinasi: ")
	        fmt.Scanln(&c)
	        fmt.Print("Masukan nilai r untuk kombinasi: ")
	        fmt.Scanln(&d)

	        fmt.Println("Permutasi: ", permutasi(a, c),"   ", "Kombinasi: ", kombinasi(a, c))
	        fmt.Println("Permutasi: ", permutasi(b, d), "    ", "Kombinasi: ", kombinasi(b, d))
    }

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](gambar/FoG.JPG)
[]

