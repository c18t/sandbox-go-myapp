package main

import "strconv"

// 引数に与えられた数値のFizzBuzzを返却します
func FizzBuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(i)
}
