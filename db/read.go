package db

import (
  "fmt"
  "io/ioutil"
  "stops/stop"

  "gopkg.in/yaml.v3"
)


func read() {
  yfile, err := ioutil.ReadFile("items.yaml")

  if err != nil {
    // create it with the future write function
  }

  data := make(map[string] stop.Stop)
  err = yaml.Unmarshal(yfile, &data)

  if err != nil {
    fmt.Println("The stops.yml file is unreadable :(");
    return
  }

  for k, v := range data {
    fmt.Printf("%s: %s\n", k, v)
  }
}
