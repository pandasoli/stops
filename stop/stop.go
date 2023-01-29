package stop

type Stop struct {
  Title string
  Shortcut string
  Description string
  When string

  Stops []*Stop
}

func New(Title string, Shortcut string, Description string, When string, Stops []*Stop) Stop {
  return Stop {
    Title,
    Shortcut,
    Description,
    When,
    Stops,
  }
}