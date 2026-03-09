package main

import "fmt"

func main () {
	var parsel, total_berat, sisa_berat, biaya_kirim int
	

	fmt.Print("Berat Parsel dalam gram : ")
	fmt.Scan(&parsel)
	
	total_berat = parsel / 1000
	sisa_berat = parsel % 1000
	fmt.Println("Detail berat : ", total_berat  ,"kg +",sisa_berat ,"gr" )
	
	

	if sisa_berat == 500 {
		fmt.Println("Detail biaya : Rp.", (total_berat * 10000), " + Rp.", (sisa_berat * 5))
		biaya_kirim =  (total_berat * 10000) + (sisa_berat * 5)
	}else if sisa_berat < 500 {
		fmt.Println("Detail biaya : Rp.", (total_berat * 10000), " + Rp.", (sisa_berat * 15))
		biaya_kirim = (total_berat * 10000) + (sisa_berat * 15)
	}else if sisa_berat > 500{
		fmt.Println("Detail biaya : Rp.", (total_berat * 10000), " + Rp.", (sisa_berat))
		biaya_kirim = (total_berat * 10000) + 0
	}	
	
	
	fmt.Println("Jadi total biaya : Rp.", biaya_kirim)


}