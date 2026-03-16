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
	var x, y, z int
	fmt.Print("Masukan nilai x, y, z: ")
	fmt.Scanln(&x,&y, &z)
	fmt.Println("Fogoh = ", (Fx(Gx(Hx(x)))))
	fmt.Println("Gohof = ", (Gx(Hx(Fx(y)))))
	fmt.Println("Hofog = ", (Hx(Fx(Gx(z)))))
}
