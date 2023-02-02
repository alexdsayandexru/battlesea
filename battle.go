package main

import "math/rand"

const (
	void   = " "
	ship   = "v"
	shadow = "."
	shot   = "*"
	kill   = "x"
)

type battlefield [][]string

type Battle struct {
	fieldInn battlefield
	fieldOut battlefield
}

func CreateBattle() *Battle {
	return &Battle{}
}

func (b *Battle) Init() {
	b.fieldInn = battlefield{
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
	}
	b.fieldOut = battlefield{
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
		{void, void, void, void, void, void, void, void, void, void},
	}
	b.genShips(10)
}

func (b *Battle) genShips(count int) {
	ships := 0
	for i := 0; i < 100; i++ {
		if ships < count && setShip(rand.Intn(10), rand.Intn(10), b.fieldInn) {
			ships++
		}
	}
}

func set(i int, j int, field battlefield, char string) {
	if i >= 0 && i < len(field) && j >= 0 && j < len(field) {
		field[i][j] = char
	}
}

func setShip(i int, j int, field battlefield) bool {
	if field[i][j] == void {
		set(i-1, j-1, field, shadow)
		set(i-1, j, field, shadow)
		set(i-1, j+1, field, shadow)
		set(i, j-1, field, shadow)
		set(i, j, field, ship)
		set(i, j+1, field, shadow)
		set(i+1, j-1, field, shadow)
		set(i+1, j, field, shadow)
		set(i+1, j+1, field, shadow)
		return true
	}
	return false
}

func (b *Battle) MakeShot(i int, j int, isPlayer bool) (bool, bool) {
	if b.fieldInn[i][j] == void {
		b.fieldInn[i][j] = shot
		b.fieldOut[i][j] = shot
		return true, false
	} else if b.fieldInn[i][j] == shadow {
		if isPlayer {
			b.fieldInn[i][j] = shot
			b.fieldOut[i][j] = shot
			return true, false
		}
		return false, false
	} else if b.fieldInn[i][j] == ship {
		b.fieldInn[i][j] = kill
		setShip(i, j, b.fieldOut)
		return true, b.isGameOver()
	}
	return false, false
}

func (b *Battle) isGameOver() bool {
	for _, row := range b.fieldInn {
		for _, cell := range row {
			if cell == ship {
				return false
			}
		}
	}
	return true
}
