package main

import (
  "net/http"
  "fmt"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := "Hello world!"

  w.Write([]byte(message))
}

func main() {
  http.HandleFunc("/", sayHello)

  fmt.Println("Listening on 8080...")  

  err := http.ListenAndServe(":8080", nil); 

  if err != nil {
    panic(err)
  } 
  
}