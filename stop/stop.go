package stop


type Stop struct {
  title string
  shortcut string
  description string
  when string

  stops []*Stop
}

func New(title string, shortcut string, description string, when string, stops []*Stop) Stop {
  return Stop {
    title,
    shortcut,
    description,
    when,
    stops,
  }
}
