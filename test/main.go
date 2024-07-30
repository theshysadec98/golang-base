package main

import (
	"fmt"
	"strconv"
)

func main() {
	var test = 123
	fmt.Println(reverse(test))
}

func reverse(x int) int {
	if x <= -(2^(31)) || x >= (2^(31)-1) {
		return 0
	}

	temp := strconv.Itoa(x)
	var sum = ""
	if temp[0] == '-' {
		sum = temp[0]
	}
	for i := 0; i < len(temp); i++ {

	}

	return int(sum)
}
