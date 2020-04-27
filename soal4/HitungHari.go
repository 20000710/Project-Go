package main

import "fmt"

func main(){
	Bulan := 30
	BulanAwal := "September"
	BulanAkhir := "September"
	TanggalAwal := 15
	TanggalAkhir := 29
	jumlah := 0

	jumlah = Bulan - TanggalAwal + TanggalAkhir

	if BulanAwal == BulanAkhir {
		jumlah = TanggalAkhir - TanggalAwal + 1
	} 	

	fmt.Printf("Jumlah hari dari tanggal %d %s hingga Tanggal %d %s adalah: %d hari", TanggalAwal, BulanAwal, TanggalAkhir, BulanAkhir, jumlah)
}