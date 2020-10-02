package main

import "math/rand"

const (
	rows    = 24
	columns = 40
)

type screen struct {
	id   int
	area [rows * columns]int
}

func genScreen(id int) *screen {
	f := screen{id: id}

	//throw down the base ground unit, depending on biome
	for i := 0; i < (rows * columns); i++ {
		f.area[i] = 1
	}
	//throw down some rocks
	for i := rand.Intn(5); i >= 0; i-- {
		f.area[rand.Intn(rows*columns)] = 2
	}
	//throw down some cacti
	for i := rand.Intn(5); i >= 0; i-- {
		f.area[rand.Intn(rows*columns)] = 3
	}

	return &f
}
