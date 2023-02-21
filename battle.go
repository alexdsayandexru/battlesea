package main

import "math/rand"

const (
	void   = " "
	ship   = "v"
	shadow = "."
	shot   = "*"
	kill   = "x"
)

type point struct {
	x int
	y int
}

type battlefield [][]string

type Battle struct {
	GameOver bool
	fieldInn battlefield
	fieldOut battlefield
}

func NewBattle() *Battle {
	return &Battle{}
}

func (b *Battle) Init(countOneShips int) {
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
	b.genShips(4, 1)
	b.genShips(3, 2)
	b.genShips(2, 3)
	b.genShips(1, 4)
}

func genVShip(countDeck int) []point {
	points := []point{}
	y := rand.Intn(10)
	x := rand.Intn(10 - countDeck + 1) + countDeck - 1
	for i:=0; i<countDeck; i++ {
		points = append(points, point{y: y, x: x - i})
	}
	return points
}

func genHShip(countDeck int) []point {
	points := []point{}
	y := rand.Intn(10 - countDeck + 1) + countDeck - 1
	x := rand.Intn(10) 
	for i:=0; i<countDeck; i++ {
		points = append(points, point{y: y - i, x: x})
	}
	return points
}

func (b *Battle) genShips(countDeck int, countShips int) bool {
	ships := 0
	for i := 0; i < 100; i++ {
		points := genVShip(countDeck)
		if placeShip(points, b.fieldInn) {
			ships++
		}
		if ships >= countShips {
			return true
		}
		points = genHShip(countDeck)
		if placeShip(points, b.fieldInn) {
			ships++
		}
		if ships >= countShips {
			return true
		}
	}
	return false
}

func thereIsVoidArea (points []point, field battlefield) bool {
	for _, p := range points {
		if field[p.y][p.x] != void {
			return false
		}
	}
	return true
}

func placeShip(points []point, field battlefield) bool {
	if thereIsVoidArea(points, field) {
		for _, p := range points {
			placeShipElement(p, field)
		}
		return true
	}
	return false
}

func placeShipElement(p point, field battlefield) {
	placeElement(p.y-1, p.x-1, field, shadow)
	placeElement(p.y-1, p.x, field, shadow)
	placeElement(p.y-1, p.x+1, field, shadow)
	placeElement(p.y, p.x-1, field, shadow)
	placeElement(p.y, p.x, field, ship)
	placeElement(p.y, p.x+1, field, shadow)
	placeElement(p.y+1, p.x-1, field, shadow)
	placeElement(p.y+1, p.x, field, shadow)
	placeElement(p.y+1, p.x+1, field, shadow)
}

func placeElement(y int, x int, field battlefield, char string) {
	if y >= 0 && y < len(field) && x >= 0 && x < len(field) && field[y][x] != ship {
		field[y][x] = char
	}
}

func (b *Battle) MakeShot(y int, x int, isPlayer bool) (bool) {
	if b.GameOver {
		return false
	}

	if (b.fieldInn[y][x] == void) || (b.fieldInn[y][x] == shadow && isPlayer) {
		b.fieldInn[y][x] = shot
		b.fieldOut[y][x] = shot
		return true
	} else if b.fieldInn[y][x] == ship {
		b.fieldInn[y][x] = kill
		placeShipElement(point{y: y, x: x}, b.fieldOut)
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
