package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
	var num1 int = 50
	var num2 float64 = 15.4
	var x string = "Hi"
	fmt.Println((x))
	fmt.Println(num1 + int(num2))

	y := 100
	fmt.Println((y))

	z := uint(0)
	fmt.Println(z)

	a := "hi"
	b := 2
	c := a + fmt.Sprint(b)
	fmt.Println(c)
	d := 3.4
	d = float64(math.Floor(d))

	e := "1234"
	f, err := strconv.Atoi(e)
	fmt.Println(f, err)

	if b < 40 {
		fmt.Println("run")
	} else if y > 40 {
		fmt.Println("stop")
	} else {
		fmt.Println("continue")
	}

	switch {
	case b < 1: // case b < 1: without switch b just switch
		fmt.Println("one")
		fallthrough
	case b == 2:
		fmt.Println("two")
	default:
		fmt.Println("default")

	}

	switch a {
	case "a", "b", "c":
		fmt.Println("a is a b or c")
	}

	for idx := 0; idx < 10; idx++ {
		fmt.Println(idx)
	}

	m := 1
	for m < 10 {
		fmt.Println("loop")
		m++
	}

	str := "Hello World!"
	val2, err2 := strconv.Atoi(string(str[0]))
	if err2 == nil {
		fmt.Println(val2)
	}

	fmt.Println(string(str[0]))

	for _, char := range str { // _ -> i
		fmt.Printf("%c", char)
	}

	for idx := 0; idx < len(str); idx++ {
		fmt.Printf("%c", str[idx])
	}

	var arr [2]int
	fmt.Println(arr)
	arr2 := [2]int{1, 2}
	arr3 := [2][2]int{{1, 2}, {3, 4}} /// ... 2
	arr3[0] = [2]int{5, 10}
	fmt.Println(arr2)
	fmt.Println(arr3)

	for i, value := range arr3 {
		fmt.Println(i, value)
		for j, value2 := range value {
			fmt.Println(j, value2)
		}
	}

	for i, value := range arr {
		fmt.Println(i, value)
	}
	test(arr3)
	fmt.Println(arr3)
	test2(&arr3)
	fmt.Println(arr3)

	arr5 := [5]int{1, 2, 3, 4, 5}
}

func test(arr [2][2]int) {
	arr[0] = [2]int{100, 100}
}

func test2(arr *[2][2]int) {
	arr[0] = [2]int{100, 100}
}

//go run demo.go

// go build demo.go
//./demo.exe
