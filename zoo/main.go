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

}
