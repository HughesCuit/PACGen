package main

import (
  "fmt"
  "net/http"
  "strings"
)

func main() {
  http.HandleFunc("/proxy/", func(w http.ResponseWriter, r *http.Request) {
    path := strings.Split(r.URL.Path, "/")
    if len(path) < 5 || path[4] != "global.pac" {
      http.Error(w, "Not found", http.StatusNotFound)
      return
    }
    host := path[2]
    port := path[3]
    js := fmt.Sprintf(`function FindProxyForURL(url, host) { return "SOCKS %s:%s"; }`, host, port)
    w.Header().Set("Content-Type", "application/x-ns-proxy-autoconfig")
    w.Write([]byte(js))
  })
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
  })
  err := http.ListenAndServe(":10086", nil)
  if err != nil {
    panic(err)
  }
}
