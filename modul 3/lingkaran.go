package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	hasil := math.Sqrt((a - c)*(a - c) + (b - d)*(b - d))
	return hasil
}

func didalam(cx, cy, r, x, y float64) bool {
	if jarak(cx, cy, x, y) < r {
		return true
	}
	return false
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Print("Nilai Lingkaran 1 cx1, cy1, r1: ")
	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Print("Nilai Lingkaran 2 cx2, cy2, r2: ")
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Print("Titik x, y: ")
	fmt.Scanln(&x, &y)

	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)
	if dalam1 && dalam2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik didalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan 2")
	}
}