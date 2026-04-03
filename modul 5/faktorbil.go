package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	faktorbil(n, 1)
}

func faktorbil(n int, i int) {
	if i > n {
		return
	}
	if n % i == 0 {
		fmt.Print(i, " ")
	}
	faktorbil(n, i + 1)
}