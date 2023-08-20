package main

type poke struct {
	name  string
	types []string
}

var pikachu = poke{
	"pikachu",
	[]string{"electric"},
}

var charmander = poke{
	"charmander",
	[]string{"fire"},
}
