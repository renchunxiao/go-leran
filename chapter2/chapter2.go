package main

import "fmt"

func main() {
	//avg := answer1([]float64{1, 2, 3, 4, 5})
	//fmt.Println(avg)
	//a, b := answer2(2, 1)
	//fmt.Println(a)
	//fmt.Println(b)
	//answer3(2, 4, 6)
	//fmt.Println(answer4(5))
	//fmt.Println(max([]int{1, 2, 3, 45}))
	//fmt.Println(min([]int{1, 2, 3, 45}))
	//n := []int{12, 2, 34, 21, 6, -1}
	//bubbleSort(n)
	//fmt.Println(n)
	//f := plusTwo()
	//fmt.Println(f(2))
	f := plusX(5)
	fmt.Println(f(2))
}

func answer1(args []float64) (avg float64) {
	if len(args) == 0 {
		return 0
	} else {
		sum := 0.0
		for _, v := range args {
			sum += v
		}
		return sum / float64(len(args))
	}
}

func answer2(a int, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func answer3(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func answer4(size int) int {
	if size == 0 {
		return 0
	} else if size == 1 {
		return 1
	} else {
		return answer4(size-1) + answer4(size-2)
	}
}

func max(l []int) (max int) {
	max = l[0]
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return
}

func min(l []int) (min int) {
	min = l[0]
	for _, v := range l {
		if v < min {
			min = v
		}
	}
	return
}

func bubbleSort(n []int) {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}

func plusTwo() func(int) int {
	return func(n int) int { return n + 2 }
}

func plusX(x int) func(int) int {
	return func(n int) int { return n + x }
}
