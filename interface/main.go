package main

import "fmt"

type saveable interface {
	save()
}

type jpeg struct {
	name   string
	width  int
	height int
}

func (j jpeg) save() {
	fmt.Printf("%s.jpeg", j.name)
}

type mp4 struct {
	name   string
	frames []int
}

func (m mp4) save() {
	fmt.Printf("%s.mp4", m.name)
}

func main() {
	myNuts := jpeg{"my ball bag", 100, 100}
	mySexTape := mp4{
		name:   "raw dogging",
		frames: []int{1, 2, 3, 4, 5},
	}

	items := []saveable{myNuts, mySexTape}
	for _, item := range items {
		item.save()
	}
}
