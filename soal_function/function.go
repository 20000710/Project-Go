package main

import "fmt"

func main() {
	status, word, ok := seeker("Samsung")

	if !ok {
		fmt.Print("status pencarian : "+status, "\nkata yang dicari : "+word+"\n")
	} else {
		fmt.Print("status pencaian : "+status, "\nkata yang dicari : "+word+"\n")
	}
}

func seeker(word string) (string, string, bool) {
	var m = [...]string{"Hp", "Asus", "Lenovo", "Acer", "Samsung"}

	for _, arr := range m {

		if arr == word {
			return "found", word, true
		}
	}

	return "not found", word, false
}
