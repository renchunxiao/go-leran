package main

import "fmt"

func main() {
	m := []int{1, 3, 4}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("%v", (Map(f, m)))
}

func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for i, v := range l {
		j[i] = f(v)
	}
	return j
}
