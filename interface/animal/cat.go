package animal

type Cat struct {
	Name  string
	Breed string
}

func (c Cat) Speak() string {
	return "Meow"
}

func (c Cat) ListensToCommands() bool {
	return false
}
