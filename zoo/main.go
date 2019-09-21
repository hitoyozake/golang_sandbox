package main

import (
	"fmt"

	"./animals" // my original package
)

var package_var = "this is package var"

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
}
