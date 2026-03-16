# <h1 align="center"> Laporan Praktikum Modul 3 </h1>
<p align="center">  [Arsal Aji Nugroho] - [109082530039] </p>

## Unguided 

### 1. [Soal]
#### soal1.go

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
![Screenshot Output Unguided 1_1](https://github.com/arsal-html/109082530039_Arsal_Aji_Nugroho/blob/main/modul_3/gambar/permutasi.JPG)
[penjelasan]
