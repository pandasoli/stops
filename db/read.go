package db

import (
  "fmt"
  "io/ioutil"
  "stops/stop"

  "gopkg.in/yaml.v3"
)


func Read() map[string]stop.Stop {
  yfile, err := ioutil.ReadFile("db/stops.yml")

  if err != nil {
    Write([]stop.Stop {})
  }

  data := make(map[string] stop.Stop)
  err = yaml.Unmarshal(yfile, &data)

  if err != nil {
    fmt.Println(err)
    return nil
  }

  return data
}

func Write(data []stop.Stop) {
  res, err := yaml.Marshal(data)

  if err != nil {
    return
  }

  ioutil.WriteFile("db/stops.yml", res, 0)
}
