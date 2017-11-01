package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Exercise that adds all divisors of 3 and 5 of 0 to until value.
Until value is received per parameter.
*/
func main() {

	input, _ := strconv.Atoi(os.Args[1])

	fmt.Println(addDivisors_3and5(input))

}

func addDivisors_3and5(until int) int {

	var add int
	for i := 0; i <= until; i++ {

		if i%3 == 0 {
			add = add + i
		}
		if i%5 == 0 {
			add = add + i
		}

	}

	return add
}
