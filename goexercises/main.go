package main

import "fmt"

func main() {

	fmt.Println(mult(4))
	fmt.Println(string_and_int("tim is cool", 50))
	fmt.Println(named_return("tim is really cool"))
}

func mult(x int) int {
	return x * x
}

func string_and_int(s string, i int) (string, int) {
	return s, i
}

func named_return(s string) (str string) {
str = s
return
}
