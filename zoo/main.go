package main

import (
	"fmt"

	"./animals" // my original package
)

var package_var = "this is package var"

func powA(a [3]int) { //copy
	for i, item := range a {
		a[i] = item * item
	}
}

func powS(a []int) { // reference
	for i, item := range a {
		a[i] = item * item
	}

}

func sub() {
	for {
		fmt.Println("Sub loop...")
	}
}

func sum(s ...int) int {
	var r int
	for _, item := range s {
		r += item
	}
	return r
}

func main() {
	fmt.Println("test!")
	fmt.Println(animals.Baw()) // using my original package
	fmt.Println(animals.MonkeyFeed())
	{
		//multiple var decleration(1)
		var x, y, z int
		x, y, z = 0, 1, 2
		fmt.Println(x, y, z)
		var n int // explicit variable
		n = 5
		fmt.Println(n)

		// multiple var decleration(2)
		var (
			x2, y2, z2 int
			name       string
		)

		x2, y2, z2, name = 1, 2, 3, "abc"
		fmt.Println(x2, y2, z2, name)
	}
	//implicit var declaration needs initialize
	{
		n := 10
		name := 100
		var n2 = 4

		var (
			x     = 1
			y     = 2
			z     = 3
			name2 = 5
		)

		fmt.Println(n, name, n2, x, y, z, name2)

	}

	// using package var
	fmt.Println(package_var)

	{
		//types
		//int8, int16, int32, int64, uint8(byte), uint1`6, uint32, uint64
		//float32, float64
		var f64 float64
		f64 = 1.2
		fmt.Println(f64)
	}

	{
		//array
		a := [4]int{1, 2, 3, 4}
		b := [4]int{1, 2} //{1,2,0,0}
		c := [4]int{}     // {0, 0, 0, 0}

		var d [4]int        //{ 0, 0, 0, 0}
		e := [...]int{1, 2} //{1, 2}

		fmt.Println(a, b, c, d, e)
		fmt.Println(b[1])
	}

	{
		//interface type
		var x interface{}
		x = 1
		x = 3.14
		x = "hello"
		x = [...]uint8{1, 2, 3}

		fmt.Println(x)

		x = 4
		i := x.(int)
		//f := x.(float64) //->error
		// type assertion

		i, isInt := x.(int)
		f, isFloat64 := x.(float64)

		fmt.Println(i, f, isInt, isFloat64)

		switch x.(type) {
		case bool:
			fmt.Println("x is bool type")
		case int, uint:
			fmt.Println("x is integer or u integer")
		default:
			fmt.Println("x is unknown type")
		}

		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}

	}

	/*
		go sub() // goroutin

		for i := 0; i < 30; i++ {

			println("main loop...")

		}
	*/

	{
		va := make([]int, 3)

		va[0] = 1
		va[1] = 2
		va[2] = 3
		fmt.Printf("%d\n", len(va))
		va = append(va, 4)
		for i, item := range va {
			fmt.Printf("index: %d, %d\n", i, item)
		}

		va2 := va[1:3] //slicing
		for i, item := range va2 {
			fmt.Printf("index: %d, %d\n", i, item)
		}

		capacity := make([]int, 8, 24) // argument of 3rd is capacity

		capacity[0] = 100

		fmt.Println(capacity[1])

		fmt.Println(sum(va...)) // variadec arguments expanded

		three := [3]int{1, 2, 3}

		powA(three)

		for _, item := range three {
			fmt.Println(item) // nochange
		}

		threeS := []int{1, 2, 3} // [...] or [NUM] -> array, [] -> slice
		powS(threeS)             // reference

		for _, item := range threeS {
			fmt.Println(item)
		}

	}

	{
		//map
		var mp map[int]string
		mp = map[int]string{1: "Taro", 2: "John"}

		mp[1] = "Taro"

		for key, value := range mp {
			fmt.Printf("K :%d, V: %s\n", key, value)
		}

	}

}
