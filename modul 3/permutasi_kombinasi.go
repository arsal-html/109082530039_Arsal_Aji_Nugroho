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