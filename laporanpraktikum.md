# <h1 align-"center">Laporan Praktikum Modul1 - 2</h1>
<p align-"center">[Arsal Aji Ngroho] - [109082530039]</p>

### 1. [SOAl]
#### temp.go

'''go
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
'''

### Output Unguided : 

##### Output