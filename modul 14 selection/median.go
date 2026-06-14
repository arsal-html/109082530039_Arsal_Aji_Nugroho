package main

import "fmt"

type arr [1000000]int

func selectionSort(T *arr, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[minIdx] {
				minIdx = j
			}
		}
		T[i], T[minIdx] = T[minIdx], T[i]
	}
}

func hitungMedian(T *arr, n int) int {
	var median int
	if n%2 != 0 {
		median = T[n/2]
	} else {
		median = (T[(n/2)-1] + T[n/2]) / 2
	}
	return median
}

func main() {
	var data arr
	var angka, n int
	fmt.Scan(&angka)
	for angka != -5313 {
		if angka == 0 {
			selectionSort(&data, n)
			fmt.Println(hitungMedian(&data, n))
		} else {
			data[n] = angka
			n = n + 1
		}
		fmt.Scan(&angka)
	}
}
