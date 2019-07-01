package main

import (
  "net"
  "bufio"
  "fmt"
  "log"
  )


func main(){
  listen, _ := net.Listen("tcp", ":10020")
  con, err := listen.Accept()

  if err != nil{
      fmt.Printf("err\n")
  }


    for {
        message, _ := bufio.NewReader(con).ReadString('\n')
        fmt.Printf("Recieved Message: %s", message)
        log.Print(message)
        con.Write([]byte("aaa"+"\n"))
  }

        con.Close()

}
