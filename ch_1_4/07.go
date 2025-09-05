package main

import "fmt"

func main() {
	var n int

	_, _ = fmt.Scan(&n)

	if n < 0 {
		fmt.Println("отрицательное число")
		return
	}

	if n == 0 {
		fmt.Println("ноль")
		return
	}

	fmt.Println("положительное число")
}
