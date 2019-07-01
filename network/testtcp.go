package main

import (
  "net"
  "bufio"
  "fmt"
)

//クライアント
func main(){
  conn, err := net.Dial("tcp", "localhost:10020")
  // connをread/write

  if err != nil{
    fmt.Printf("error client")
  }

  text := "test"
  conn.Write([]byte(text+"\n"))

  message, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Printf("%s", message)

  text = "q"
  conn.Write([]byte(text+"\n"))

  message, _ = bufio.NewReader(conn).ReadString('\n')
  fmt.Printf("%s", message)


}
