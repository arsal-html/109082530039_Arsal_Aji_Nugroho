package main

import "fmt"

type isiArr [100000]int

func selectionSort(T *isiArr, n int) {
	var idx_min int

	i := 1
	for i <= n-1 {

		idx_min = i - 1

		j := i
		for j < n {

			if T[idx_min] > T[j] {
				idx_min = j
			}

			j = j + 1
		}

		t := T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t

		i = i + 1
	}
}

func main() {

	var n, m int
	var a isiArr

	fmt.Scan(&n)

	i := 0

	for i < n {

		fmt.Scan(&m)

		j := 0
		for j < m {

			fmt.Scan(&a[j])
			j++
		}

		selectionSort(&a, m)

		j = 0
		for j < m {

			if a[j]%2 != 0 {
				fmt.Print(a[j], " ")
			}

			j++
		}

		j = m - 1

		for j >= 0 {

			if a[j]%2 == 0 {
				fmt.Print(a[j], " ")
			}

			j--
		}

		fmt.Println()

		i++
	}
}