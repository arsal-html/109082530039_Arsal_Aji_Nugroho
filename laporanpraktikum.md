# <h1>Laporan Praktikum Modul 1 - 2</h1>
<p align-"center">[Arsal Aji Ngroho] - [109082530039]</p>

### 1. [SOAl]
#### Soal1.go

'''go
package main

import "fmt"

func main() {

var (
satu, dua, tiga string
<br>temp string
)

fmt.Print("Masukan input string: ")
<br>fmt.Scanln(&satu)
<br>fmt.Print("Masukan input string: ")
<br>fmt.Scanln(&dua)
<br>fmt.Print("Masukan input string: ")
<br>fmt.Scanln(&tiga)
<br>fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
<br>temp = satu
<br>satu = dua
<br>dua = tiga
<br>tiga = temp
<br>fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
'''

### Output Unguided : 

##### Output
![Screenshot Output Unguided 1_1](https://github.com/arsal-html/109082530039_Arsal_Aji_Nugroho/ALPRO_2/modul_1_dan_2/soal1-PNG)
[]