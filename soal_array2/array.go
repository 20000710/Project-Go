package main

import "fmt"

func main()  {
	arr := [30]int{3,7,9,12,14,5,6,10,13,11,15,18,16,2,1,19,17,20,4,8}	

	var genap [] int
	var ganjil [] int
	var prima [] int

	for i := 0; i < len(arr); i++{
		if arr[i] % 2 == 0 {
			genap = append(genap, arr[i])
		} else {
			ganjil = append(ganjil, arr[i])
		}
		
	}

	for a := 1; a < len(arr); a++ {
		if arr[a]%2 != 0 && arr[a] > 1 {
			prima = append(prima, arr[a])
		}
	}

	fmt.Println("Jumlah bilangan genap", genap, "jumlah data genap", len(genap))
	fmt.Println("Jumlah bilangan ganjil", ganjil, "jumlah data ganjil", len(ganjil))
	fmt.Println("jumlah bilangan prima", prima, "jumlah data prima", len(prima))
}