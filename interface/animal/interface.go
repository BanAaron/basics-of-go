package animal

// CanSpeak is an interface. The methods inside it are required for the type
type CanSpeak interface {
	// Speak and ListensToCommands are both required
	Speak() string
	ListensToCommands() bool
}
