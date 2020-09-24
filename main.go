package main

import (
  "encoding/json"
  "net/http"
  "time"
)

type response struct {
  Time string `json:"time"`
}

func apiTime(w http.ResponseWriter, req *http.Request) {
  res := response{time.Now().Format(time.RFC3339)}
  json, err := json.Marshal(res)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(json)
}

func main() {
  http.HandleFunc("/time", apiTime)
  http.HandleFunc("/", mainPage)
  http.ListenAndServe(":8795", nil)
}

func mainPage(res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("It is my first GO web page"))
}
