package main

import "fmt"

func main(){

    //var a, b, c int
    //var x = 1
    var y = 14
    //var(
      //  y, z int
      //  name string
    //)

    fmt.Printf("%d\n", y )

    for i:= 0; i < 5; i++ {
        fmt.Printf("%d\n", i )
    }
    for i:= 4; i > 0; i-- {
        fmt.Printf("%d\n", i )
    }

    //初期化する場合は varをつけない
    sarray := [3] string{"apple", "orange", "peach"}

    for index, s := range sarray {
        fmt.Printf("%d %s", index, s)
    }

}
