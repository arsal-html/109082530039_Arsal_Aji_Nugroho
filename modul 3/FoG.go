package main

import "fmt"

func Fx(x int) int {
	hasil := x * x
	return hasil
}

func Gx(x int) int {
	hasil := x - 2
	return hasil
}

func Hx(x int) int {
	hasil := x + 1
	return hasil
}

func main() {
	var a, b, c  int
	fmt.Print("Masukan nilai a, b, c: ")
	fmt.Scanln(&a,&b, &c)
	fmt.Println("Fogoh = ", (Fx(Gx(Hx(a)))))
	fmt.Println("Gohof = ", (Gx(Hx(Fx(b)))))
	fmt.Println("Hofog = ", (Hx(Fx(Gx(c)))))
}
