package main

import "fmt"
import "unicode/utf8"

func main() {
	//	answer1()
	//	answer2()
	//      answer3()
	answer4()
}

func answer1() {
	fmt.Println("answer1 begin ...")
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	a := 1
Here:
	fmt.Println(a)
	a++
	if a < 10 {
		goto Here
	}
	arr := []int{3, 43, 23, 4, 1}
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("answer1 end ...")
}

func answer2() {
	fmt.Println("answer2 begin ...")
	for i := 1; i <= 100; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("answer2 end ...")
}

func answer3() {
	fmt.Println("answer3 begin ...")
	str := "A"
	for i := 1; i <= 10; i++ {
		fmt.Println(str)
		str = str + "A"
	}
	langstr := "asd"
	fmt.Println(len([]byte(langstr)))
	fmt.Println(utf8.RuneCount([]byte(langstr)))
	s := "thisanewlinea"
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Println(string(r))
	s2 := "foobar"
	r2 := []rune(s2)
	for i, j := 0, len(s2)-1; i < j; i, j = i+1, j-1 {
		r2[i], r2[j] = r2[j], r2[i]
	}
	fmt.Println(string(r2))
	fmt.Println("answer3 end ...")
}

func answer4() {
	fmt.Println("answer4 begin ...")
	arr := []float64{1, 6, 2, 5}
	xs := arr[:]
	sum := 0.0
	if len(xs) == 0 {
		fmt.Println(0)
	} else {
		for _, v := range xs {
			sum = sum + v
		}
		fmt.Println(sum / float64(len(xs)))
	}
	fmt.Println("answer4 end ...")
}
