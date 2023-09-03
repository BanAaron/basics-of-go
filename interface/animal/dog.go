package animal

type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Speak() string {
	return "Woof"
}

func (d Dog) ListensToCommands() bool {
	return true
}
