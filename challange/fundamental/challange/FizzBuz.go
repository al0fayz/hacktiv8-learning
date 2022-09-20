package challange

import "fmt"

func FizzBuz() {
	var n = 15
	var i = 1
	for i <= n {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
		i++
	}
}
