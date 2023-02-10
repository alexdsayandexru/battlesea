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
	GameOver bool
	fieldInn battlefield
	fieldOut battlefield
}

func NewBattle() *Battle {
	return &Battle{}
}

func (b *Battle) Init(countShips int) {
	b.GameOver = false
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
	//b.genShips(countShips)
	b.genShipsTest()
}

func (b *Battle) genShips(countShips int) {
	ships := 0
	for i := 0; i < 100; i++ {
		if ships < countShips && setShip(rand.Intn(10), rand.Intn(10), b.fieldInn) {
			ships++
		}
	}
}

func (b *Battle) genShipsTest() {
	setShip(1, 7, b.fieldInn)
	setShip(7, 9, b.fieldInn)
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

func (b *Battle) MakeShot(i int, j int, isPlayer bool) (bool) {
	if b.GameOver {
		return false
	}

	if (b.fieldInn[i][j] == void) || (b.fieldInn[i][j] == shadow && isPlayer) {
		b.fieldInn[i][j] = shot
		b.fieldOut[i][j] = shot
		return true
	} else if b.fieldInn[i][j] == ship {
		b.fieldInn[i][j] = kill
		setShip(i, j, b.fieldOut)
		b.GameOver = b.isGameOver()
		return true
	}
	return false
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
