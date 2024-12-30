package main

import "fmt"

func product(factors ...int) (p int) {
	p = 1

	for _, factor := range factors {
		p *= factor

	}

	return

}

func main() {
	fmt.Println(product(1, 2, 3, 4, 5))
	fmt.Println(product(1, 2, 3, 4, 5, 6))
}
