package main

import "fmt"

func main(){
	Bulan := 30
	TanggalAwal := 15
	TanggalAkhir := 4
	jumlah := Bulan - TanggalAwal + TanggalAkhir

	fmt.Printf("Jumlah hari dari tanggal %d April hingga Tanggal %d adalah: %d hari", TanggalAwal, TanggalAkhir, jumlah)
}