package main

import "fmt"

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukan nilai: ")
	fmt.Scanln(&a, &b, &c, &d)

	fmt.Println("Permutasi ke-1: ", permutasi(a, c),"   ", "Kombinasi ke-1: ", kombinasi(a, c))
	fmt.Println("Permutasi ke-2: ", permutasi(b, d), "    ", "Kombinasi ke-2: ", kombinasi(b, d))
}