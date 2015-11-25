package main

import (
  "io"
  "net/http"
  "bindata"
  //"github.com/davecgh/go-spew/spew"
)

func hello(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  io.WriteString(w, "<h3>R2</h3>")
  io.WriteString(w, "<img src='/images/r2.jpg'/>")
  io.WriteString(w, "<p>A Slackbot Droid.</p>")
  io.WriteString(w, "<p>point your outbound webhook to <a href='/webhook'>this server</a></p>")
}

func r2jpg(w http.ResponseWriter, r *http.Request) {
  //w.Header().Set("Content-Type", "image/jpeg")
  data, _ := bindata.Asset("data/r2.jpg") 
  //spew.Dump( data )
  w.Write( data )
}

func main() {
  http.HandleFunc("/", hello)
  http.HandleFunc("/images/r2.jpg", r2jpg)

  http.ListenAndServe(":8000", nil)
}

