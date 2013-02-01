package boop

import (
  "fmt"
  "net/http"
  "os"
  "os/exec"
)

var commandMap map[string]string

func BoopMain() {
  loadCommandMap()
  fmt.Println("Starting http server on port 5090")
  http.Handle("/", http.HandlerFunc(httpRequestHandler))
  err := http.ListenAndServe("0.0.0.0:5090", nil)
  fmt.Println("Listening...")
  if err != nil {
    fmt.Println("ListenAndServ Error: ", err)
  }
}

func loadCommandMap() {
  commandMap = map[string]string{}
  commandMap["/test"] = "touch WOAH_MAN"
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
  fmt.Println("Received request for " + req.URL.Path)
  if v, ok := commandMap[req.URL.Path]; ok {
    fmt.Println("This corresponds with command " + v)
    fmt.Println("Executing...")
    cmd := exec.Command(os.Getenv("SHELL"), "-c", v) 
    err := cmd.Run()
    if err != nil {
      fmt.Println("Execution returned an error.  Sending status HTTP 500 (Internal Server Error)")
      w.WriteHeader(500)
      fmt.Fprint(w, "500 Internal Server Error")
    } else {
      fmt.Println("Execution returned no error.  Sending status HTTP 200 (OK)")
      w.WriteHeader(200)
      fmt.Fprint(w, "200 OK")
    }
  } else {
    fmt.Println("No command associated with this path.  Sending status HTTP 404 (Not Found)")
    w.WriteHeader(404)
    fmt.Fprint(w, "404 Not Found")
  }
}
