package main

import "fmt"
import "strconv"

type stack struct {
	index int
	data  [10]int
}

func main() {
	var s stack
	s.push(1)
	s.push(2)
	fmt.Println(s)
}

func (s *stack) push(v int) {
	s.data[s.index] = v
	s.index++
}

func (s *stack) pop() int {
	s.index--
	return s.data[s.index]
}

func (s stack) String() string {
	var str string
	for i := 0; i <= s.index; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
