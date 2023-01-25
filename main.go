package main

import (
  "fmt"
  "stops/db"
)

func main() {
  fmt.Println("\033[1;37mAdd a new Stop\033[0m")

  db.read()
}
