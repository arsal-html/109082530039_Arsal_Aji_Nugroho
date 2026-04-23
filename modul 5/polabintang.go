package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	pola(n)
}

func cetakBintang(n int) {
    if n == 0 { 
		return 
	}
    fmt.Print("*")
    cetakBintang(n - 1)
}

func pola(n int) {
    if n == 0 { 
		return 
	}
    pola(n - 1)
    cetakBintang(n)
    fmt.Println()
}