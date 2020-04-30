package main

import "fmt"

//Sepeda ...
type Sepeda struct {
	jumlahBan  int
	jumlahGigi int
	warna 		string
}

func (s *Sepeda) waktuTempuh(jarak float32) float32 {
	return jarak * 2.5 
}

func main() {	
	sepeda := []Sepeda{}

	s1 := Sepeda{jumlahBan: 2, jumlahGigi: 1, warna: "hitam"}
	s2 := Sepeda{jumlahBan: 2, jumlahGigi: 4, warna: "merah"}
	s3 := Sepeda{jumlahBan: 2, jumlahGigi: 3, warna: "hijau"}
	s4 := Sepeda{jumlahBan: 2, jumlahGigi: 5, warna: "biru cerah"}
	s5 := Sepeda{jumlahBan: 2, jumlahGigi: 9, warna: "kuning"}

	sepeda = append(sepeda, s1, s2, s3, s4, s5)

	for _, n := range sepeda {
		fmt.Println("Jumlah ban:", n.jumlahBan, " Jumlah gigi:", n.jumlahGigi, " warna:", n.warna)
		fmt.Print("Waktu tempuh sepeda :", n.waktuTempuh(19), " menit\n\n")
	}

}
