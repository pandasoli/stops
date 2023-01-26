package main

import (
  "fmt"
  "stops/db"
)

func main() {
  fmt.Println("\033[1;37mAdd a new Stop\033[0m")
  data := db.Read()

  for _, v := range data {
    fmt.Printf("\033[1;37m%s\033[0m\n", v.title)
  }
}
