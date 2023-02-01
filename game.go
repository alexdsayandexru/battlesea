package main

import "math/rand"

type battlefield [][]string

type Battle struct {
	fieldInn battlefield
	fieldOut battlefield
}

func CreateBattle () *Battle {
	return &Battle{}
}

func (b *Battle) Init() {
	b.fieldInn = battlefield{
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
	}
	b.fieldOut = battlefield{
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "," "," "},
	}

	ships := 0
	for i:=0; i<100; i++ {
		if ships < 5 && setShip(rand.Intn(10),rand.Intn(10), b.fieldInn) {
			ships++
		}
	}
	
}

func set(i int, j int, field battlefield, ch string) {
	if i >= 0 && i < len(field) && j >= 0 && j < len(field) {
		field[i][j] = ch
	}
}

func setShip(i int, j int, field battlefield) bool {
	ch := "."
	if field[i][j] == " " {
		set(i-1, j-1, field, ch)
		set(i-1, j, field, ch)
		set(i-1, j+1, field, ch)
		set(i, j-1, field, ch)
		set(i, j, field, "v")
		set(i, j+1, field, ch)
		set(i+1, j-1, field, ch)
		set(i+1, j, field, ch)
		set(i+1, j+1, field, ch)
		return true
	}
	return false
}

func (b *Battle) shot(i int, j int) (bool, bool) {
	if b.fieldInn[i][j] == " " || b.fieldInn[i][j] == "." {
		b.fieldInn[i][j] = "*"
		b.fieldOut[i][j] = "*"
		return true, false
	} else if  b.fieldInn[i][j] == "v" {
		b.fieldInn[i][j] = "x"
		setShip(i,j,b.fieldOut)
		return true, b.isGameOver()
	}
	return false, false
}

func (b *Battle) isGameOver() bool {
	ship := 0
	for _, row := range b.fieldInn {
		for _, e := range row {
			if e == "v" {
				ship++
			}
		}
	}
	return ship == 0
}