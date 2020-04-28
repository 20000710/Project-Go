package main

import "fmt"

func main() {
	s := make([]int, 0, 20)
	y := 20

	for x := 1; x < y; x++ {	
		d := 0

		for i := 1; i < y; i++ {

			if x%i == 0 {
				d++
			}				
		}

		if d == 2 && x > 1 {
			s = append(s, x)
		}
	}

	fmt.Println(s)
}