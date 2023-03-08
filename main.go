package main

import (
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello World!")
  })
  r.GET("/proxy/:host/:port/global.pac", func(c *gin.Context) {
    host := c.Param("host")
    port := c.Param("port")
    jsContent := fmt.Sprintf(`function FindProxyForURL(url, host) { return "SOCKS %s:%s"; }`, host, port)
    c.String(http.StatusOK, jsContent)
  })
  r.Run(":10086")
}
