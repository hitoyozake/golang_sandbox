package main

import (
    "fmt"
)

type Square interface{
    getSquare() int
    print_s()
}

type Circle struct{
    x int
    y int
    r int
}

func (c Circle) getSquare() int{
    return c.r*c.r*3
}

func (c Circle) print_s(){
    fmt.Println(c.r, c.x, c.y)
}


func main(){
    circle := Circle{ 0, 0, 4 }

    inf := Square(circle)
    fmt.Println(circle.r)
    inf.print_s()
}




