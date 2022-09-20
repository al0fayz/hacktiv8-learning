package challange

import (
	"fmt"
)

func Calculating() {
	var greeting = "Selamat malam"
	mapstring := map[string]int{}
	for index, char := range greeting {
		fmt.Println(string(char))
		mapstring[string(char)] = index
	}
	fmt.Println(mapstring)
}
