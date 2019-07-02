package main

import (
  "fmt"
  "net/http"
  "log"
  "bytes"
  "net/url"
)


func main(){
  client := &http.Client{}

  encoded, err := url.QueryUnescape(`{"text":"post from golang", "username": "testBot"}`)
  buf := []byte(encoded)

  request, err := http.NewRequest("POST",
    "https://hooks.slack.com/services/T8JTRHP1B/B8JQKM3AQ/*******",
     bytes.NewBuffer(buf))

  request.Header.Set("Content-Type", "application/json")

  resp, err := client.Do(request)

  if err != nil{
    log.Printf("error!")
    panic(err)
  }

  fmt.Print(resp)
  defer resp.Body.Close()


}
