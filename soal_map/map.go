package main

import "fmt"

type barang struct {
	nama  string
	harga int
	stok  int
}

func main() {
	var produk = make(map[int]barang)

	produk[1] = barang{nama: "pasta gigi", harga: 17000, stok: 13}
	produk[2] = barang{nama: "soft drink", harga: 5000, stok: 8}
	produk[3] = barang{nama: "cookies cake", harga: 12000, stok: 9}
	produk[4] = barang{nama: "shampo", harga: 2500, stok: 17}
	produk[5] = barang{nama: "mie instan", harga: 2500, stok: 10}

	fmt.Print("Produk yang memiliki stok dibawah 10: \n\n")

	for _, v := range produk {
		if v.stok < 10 {
			fmt.Println("- " + v.nama)
		}
	}
}
