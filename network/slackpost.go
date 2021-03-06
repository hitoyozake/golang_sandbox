package main

import (
  "fmt"
  "net/http"
  "log"
  "bytes"
  "net/url"
)

type SlackClient struct{
    url string
    clien *http.Client
}

//func post(client SlackClient){
//}


func main(){
  client := &http.Client{}

  var cl = &SlackClient{
    url: "abc",
    clien: &http.Client{},
  }

  encoded, _ := url.QueryUnescape(`{"text":"post from golang", "username": "testBot"}`)
  buf := []byte(encoded)

  request, _ := http.NewRequest("POST",
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
