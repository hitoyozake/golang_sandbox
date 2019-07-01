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
  reader := bufio.NewReader(con)
  for {
      fmt.Printf("waiting")
      message, err := reader.ReadString('\n')

      if err != nil{
        fmt.Printf("error!")
        return
      }

      fmt.Printf("Recieved Message: %s", message)
      fmt.Printf("")
      log.Print(message)
      con.Write([]byte(message+"\n"))

      if message == "q\n"{
        break
      }

  }

  con.Close()

}
