package main

import (
	"fmt"
)

type ints []int

func main() {
	mySLice := ints{}

	for i := 0; i < 11; i++ {
		mySLice = append(mySLice, i)
	}

	for i := 0; i < len(mySLice); i++ {
		if mySLice[i]%2 == 0 {
			fmt.Println(mySLice[i], "is even")
		} else {
			fmt.Println(mySLice[i], "is odd")
		}
	}

	fmt.Println(mySLice)
}
