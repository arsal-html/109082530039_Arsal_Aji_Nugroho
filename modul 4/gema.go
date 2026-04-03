package main

	import "fmt"

	func hitungSkor(soal *int, skor *int) {
		var waktu int
		*soal = 0
		*skor = 0
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu)
			if waktu <= 300 {
				*soal++
				*skor += waktu
			}
		}
	}
	
	func main() {
    var nama, pemenang string
    var soal, skor, jmlhsoal, jmlhskor int
    jmlhskor = 999999

	fmt.Scan(&nama)

	for nama != "selesai" {
		hitungSkor(&soal, &skor)
		
        if soal > jmlhsoal || (soal == jmlhsoal && skor < jmlhskor) {
            jmlhsoal = soal
            jmlhskor = skor
            pemenang = nama
    		}
		fmt.Scan(&nama)
	}
		
	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, jmlhsoal, jmlhskor)
		}
	}
