package main

import "fmt"

func Sum(a, b, c, d, e int) int {
	return a+b+c+d+e
}

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}
