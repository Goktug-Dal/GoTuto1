package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
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
	sl := arr5[1:3]
	sl[0] = 100
	fmt.Println(sl) // 2,3
	//:2 first 2
	//1: every element after first one

	fmt.Println(sl, len(sl), cap(sl))

	sl2 := []string{"Hello World"}
	fmt.Println(sl2)

	for x := 0; x < 10; x++ {
		sl2 = append(sl2, "tim")
		fmt.Println(sl2, len(sl2), cap(sl2)) // cap doubles if it reaches limit, len increases one per word
	}

	sl3 := make([]int, 10)
	fmt.Println(sl3)

	sl4 := []string{"Hello", "World", "Hi"}
	for i, value := range sl4 {
		fmt.Println(i, value)
	}

	test4(sl4) // it works
	fmt.Println(sl4)

	var mp map[string]int = map[string]int{"a": 1}
	for i, value := range mp {
		fmt.Println(i, value)
	}

	mp2 := map[string]int{"b": 4}
	fmt.Println(mp2)
	mp3 := make(map[string]int)
	fmt.Println((mp3))
	mp3["c"] = 1
	mp3["d"] = 2

	value, ok := mp["d"]
	fmt.Println(value, ok)
	delete(mp3, "c")

	val := add(5, 4)
	vals1, strs1 := add2(4, 5)
	fmt.Println(vals1, strs1)
	fmt.Println(val)

	val10 := callFund(doubleNum)
	fmt.Println(val10)

	val11 := callFund(func(x int) int {
		return x + 11
	})

	fmt.Println(val11)
	s := sums([]int{1, 2, 5, 657, 7, 8}...)
	fmt.Println(s)
	fmt.Println(sums(1, 2, 4, 4, 6, 6, 7))

	p1 := Person{"Tim", 23} // name: age:
	p1.name = "John"
	fmt.Println(p1)
	fmt.Println(getName(p1))

	p1.setName("Jon")

	p2 := Person2{"Alex", 20, []Sport{{"Soccer", "D"}}}
	fmt.Println(p2.favSports[0].position)

	var s11 Shape = Triangle{1, 3, 4}
	fmt.Println(s11.getPerimeter())

	defer deferredFunc()
	//panic("This is panic")
	fmt.Println("run")

	result, err2 := div(1, 0)
	if err2 != nil {
		fmt.Println("error occurred")
	} else {
		fmt.Println(result)
	}

	valss := addy(4.3, 5.6)
	fmt.Println(valss)

	X1 := 0
	Y1 := &X1
	fmt.Println(X1, Y1)

	addy2(&X1)
	fmt.Println(X1)

	go run() // if you use just run() 's it will take llonger because it wont be async
	go run2()
	go run3()
	time.Sleep(7 * time.Second)
	fmt.Println("Done")

	ch := make(chan int)
	go add110(1, 1, ch)

	x110 := <-ch
	fmt.Println(x110)

	add110(5, 100, ch)
	add110(5, -100, ch)
	add110(-3, 0, ch)
	add110(0, 4, ch)
	x155 := <-ch
	x155 = <-ch
	x155 = <-ch
	x155 = <-ch
	fmt.Println(x155)

	ch2 := make(chan int)
	ch3 := make(chan int)

	go add1109(10, 5, ch2, 4)
	go add1109(20, 15, ch3, 2)

	for i := 0; i < 2; i++ {
		select {
		case m := <-ch2:
			fmt.Println(m)
		case m2 := <-ch3:
			fmt.Println(m2)
		}
	}
}

func add1109(x int, y int, ch chan int, value int) {
	time.Sleep(time.Duration(value) * time.Second)
	ch <- x + y // channel
}

func add110(x int, y int, ch chan int) {
	ch <- x + y // channel
}

func run() {
	time.Sleep(2 * time.Second)
	fmt.Println("run")
}

func run2() {
	time.Sleep(4 * time.Second)
	fmt.Println("run2")
}

func run3() {
	time.Sleep(6 * time.Second)
	fmt.Println("run3")
}

func addy2(x *int) {
	*x += 5
}

func addy[T int | float64](x T, y T) T {
	return x + y
}

func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by 0")
	}
	return a / b, nil
}
func deferredFunc() {
	fmt.Println("defer")
	r := recover()
	if r == nil {
		fmt.Println("no error")
	} else {
		fmt.Println(r)
	}
}

func (t Triangle) getPerimeter() uint {
	return t.a + t.b + t.c
}

type Triangle struct {
	a uint
	b uint
	c uint
}

type Shape interface {
	getPerimeter() uint
}

type Sport struct {
	name     string
	position string
}

type Person2 struct {
	name      string
	age       uint
	favSports []Sport
}

func (p Person) setName(newName string) {
	//p.name = newName
}

func getName(p Person) string {
	return p.name
}

type Person struct {
	name string // lower  case = private, first letter capital = public same for function names
	age  uint
}

func sums(nums ...int) int {
	s := 0
	for _, value := range nums {
		s += value
	}
	return s
}

func callFund(callable func(int) int) int {
	return callable(10)
}

func doubleNum(num int) int {
	return 2 * num
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func add2(num1 int, num2 int) (int, string) {
	return num1 + num2, "Hello"
}

func test4(arr []string) {
	arr[0] = "change this"
}

func test(arr [2][2]int) {
	//arr[0] = [2]int{100, 100}
}

func test2(arr *[2][2]int) {
	arr[0] = [2]int{100, 100}
}

//go run demo.go

// go build demo.go
//./demo.exe
