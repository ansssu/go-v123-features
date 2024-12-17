package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println("\n\n")
	iteratorOverValues(s)

	fmt.Println("\n\n")
	iteratorOverIndexesAndValues(s)

	fmt.Println("\n\n")
	iteratorSliceBackward(s)

	fmt.Println("\n\n")
	iteratorCollects(s)

}

func iteratorOverValues(s []string) {
	fmt.Println("****** iterator over values ******")
	for v := range slices.Values(s) {
		fmt.Println(v)
	}
}

func iteratorOverIndexesAndValues(s []string) {
	fmt.Println("****** iterator over key/value pair ******")
	for i, v := range slices.All(s) {
		fmt.Println(i, v)
	}
}

func iteratorSliceBackward(s []string) {
	fmt.Println("****** iterator over key/value pair, backward the slice ******")
	for i, v := range slices.Backward(s) {
		fmt.Println(i, v)
	}
}

func iteratorCollects(s []string) {
	fmt.Println("****** iterator collects values fro iterator into a new slice ******")
	s2 := slices.Collect(slices.Values(s))

	for v := range slices.Values(s2) {
		fmt.Println(v)
	}

}
