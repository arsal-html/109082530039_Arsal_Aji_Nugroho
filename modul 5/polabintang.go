package main

import "fmt"

func main() {
	var n int
	 fmt.Scan(&n)
	polabintang(n)
}

func polabintang(n int) {
	if n <= 0 {
		return 
		}
	polabintang(n - 1)

		for i := 0; i < n; i++ {
		fmt.Print("* ")
		}
	fmt.Println()
	}