package main

import "fmt"

func main() {
	s := []string{"hello", "world"}
	for i, x := range Backwards(s) {
		fmt.Println(i, x)
	}
}

func Backwards(s []string) func(func(int, string) bool) {
	return func(yeld func(int, string) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yeld(i, s[i]) {
				return
			}
		}
	}
}
