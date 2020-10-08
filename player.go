package main

import "math/rand"

//player related functions

const (
	WIZARD   = iota
	KNIGHT   = iota
	HUNTSMAN = iota
	SAMURAI  = iota
)

const (
	HUMAN = iota
	FAIRY = iota
	ELF   = iota
	DWARF = iota
)

//maybe some sort of table to store pertinent information and functions on roles later...
//this is kind of crude...
//it is possible that the entire stats and starting traits for each role could be stored as const
//values and a series of functions
func roletoText(role int) string {
	if role == WIZARD {
		return "WIZARD"
	} else if role == KNIGHT {
		return "KNIGHT"
	} else if role == HUNTSMAN {
		return "HUNTSMAN"
	} else if role == SAMURAI {
		return "SAMURAI"
	}
	return "ONION"
}

type character struct {
	name         string
	exp          int
	role         int
	creature     int
	strength     int
	dexterity    int
	constitution int
	intelligence int
	wisdom       int
	charisma     int
	xpos         int
	ypos         int
	inventory    *int
}

func newChar(name string, role int, creature int) *character {
	c := character{name: name}
	c.role = role
	c.creature = creature
	return &c
}

func genStats(c *character) {
	//also add the creature and role traits.
	c.strength = rand.Intn(4) + 8
	c.dexterity = rand.Intn(4) + 8
	c.constitution = rand.Intn(4) + 8
	c.intelligence = rand.Intn(4) + 8
	c.wisdom = rand.Intn(4) + 8
	c.charisma = rand.Intn(4) + 8
}
