package dog

type Dog struct {
	id    int
	name  string
	noise string
}

func (d Dog) Bark() string {
	return d.noise
}

// NewDog is a factory. It is just a function that returns Dog
func NewDog(name string, noise string) Dog {
	return Dog{id: 1, name: name, noise: noise}
}
