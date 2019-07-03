package main

import (
  "fmt"
)

func getFunc() (int, int, int) {
  return 1, 2, 3
}

func main(){
    s := "abc"

    fmt.Printf("%s", s)
    a, b, _ := getFunc()

    fmt.Printf("%d %d", a, b)


    ax, b, c := getFunc()

    fmt.Printf("%d %d %d", ax, b, c)

}
