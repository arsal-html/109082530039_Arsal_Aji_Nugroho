package main

import "fmt"

func angkaMembesar(T [1000]int, n int) {
	i := 1
	for i <= n-1 {
		j := i
		temp := T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()

	 if n <= 1 {
		fmt.Print()
	 }else {
	jarak := T[1] - T[0]
	konsisten := true
	for i := 1; i < n-1; i++ {
    if T[i+1] - T[i] != jarak {
        konsisten = false
        break
    	}
	}
	if konsisten {
		fmt.Print("Data berjarak ", jarak)
	}else {
		fmt.Print("Data berjarak tidak tetap")
		}
	}
}

func main() {
	var T [1000]int
	var x int
	n := 0
 
	fmt.Print("Masukkan Angka: ")
	for {
		fmt.Scan(&x)
		if x < 0 {
			break 
		}
		T[n] = x
		n++
	}

	angkaMembesar(T, n)
}