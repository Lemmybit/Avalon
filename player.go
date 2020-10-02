package main

//player related functions

const (
	barbarian = iota
	wizard    = iota
	rogue     = iota
	ranger    = iota
)

type player struct {
	name      string
	class     int
	xpos      int
	ypos      int
	inventory [10]int
}

func newPlayer(name string, class int) *player {
	p := player{name: name, class: class}
	p.xpos = 0
	p.ypos = 0
	p.inventory[1] = 1 //set to an arbitrary item number
	return &p
}
