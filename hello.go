package main

import "fmt"

func div60(i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in div60", r)
		}
	}()

	fmt.Println(60/i)
}

func main() {
	arr := []int{1, 2, 3, 0, 4, 5}
	for _, v := range arr {
		div60(v)
	}
}