package main

import (
  "fmt"
  "net/http"
  "os"
)

var commandMap map[string]string

func main() {
  loadCommandMap()
  fmt.Println("Starting http server on port 5090")
  http.Handle("/", http.HandlerFunc(httpRequestHandler))
  err := http.ListenAndServe("0.0.0.0:5090", nil)
  if err != nil {
    fmt.Println("ListenAndServ Error: ", err)
  }
}

func loadCommandMap() {
  commandMap["/test"] = "touch WOAH_MAN"
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
}
