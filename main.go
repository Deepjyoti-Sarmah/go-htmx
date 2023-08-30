package main

import (
	"fmt"
	"net/http"
)

func main() {
  fmt.Println("hello world")

  http.ListenAndServe(":8000", nil)
}
