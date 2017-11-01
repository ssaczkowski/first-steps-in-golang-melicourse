package main

import (
	"fmt"
	"os"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

/*
Program that receives a list of integers per parameter and stores them in a slice.
Then, populate a binary tree and print in ascending order.
*/
func main() {

	abb := new(tree)

	for i := 0; i < len(os.Args); i++ {
		input, err := strconv.Atoi(os.Args[i])
		if err == nil {
			add(abb, input)
		}
	}

	show(abb)

}

func add(t *tree, value int) *tree {

	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func show(t *tree) {

	if t != nil {
		show(t.left)
		fmt.Printf("Value: %d \n", t.value)
		show(t.right)
	}

}
