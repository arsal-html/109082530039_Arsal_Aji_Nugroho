package main

import "fmt"

func main() {
	var a, b, c, d string
	
	hasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("percobaan ", i, ": ")
		fmt.Scan(&a, &b, &c, &d)
		if a != "merah" || b != "kuning" || c != "hijau" || d != "ungu" {
			hasil = false
		}
	}
	fmt.Println(hasil)
}
