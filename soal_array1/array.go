package main

import "fmt"

func main() {
	arr := [20]int{2, 1, 3, 5, 7, 3, 6, 9, 10, 8, 13, 11, 14, 12, 16, 18, 15, 17, 20, 19}
	arr2 := [20]int{17, 22, 15, 25, 4, 9, 10, 12, 11, 21, 19, 17, 14, 18, 16, 3, 13, 7, 1, 2}

	fmt.Printf("Angka yang tidak sama antara arr dan arr2 : ")
	fmt.Printf("%+v\n", seperator(arr, arr2))
}

func seperator(arr [20]int, arr2 [20]int) []int {
	var diff []int

	for i := 0; i < 2; i++ {

		for _, i := range arr {
			f := false

			for _, j := range arr2 {
				if i == j {
					f = true
					break
				}

			}

			if !f {
				diff = append(diff, i)
			}
		}

		if i == 0 {
			arr, arr2 = arr2, arr
		}
	}

	return diff

}
