package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Game struct {
	comp   *Battle
	player *Battle
}

func CreateGame() *Game {
	game := Game{comp: CreateBattle(), player: CreateBattle()}
	game.comp.Init()
	game.player.Init()
	return &game
}

func (g *Game) Print() {
	header := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Print(" ")
	fmt.Print(header)
	fmt.Println(header)
	for i := 0; i < len(g.comp.fieldOut); i++ {
		fmt.Print(i)
		fmt.Print(g.comp.fieldOut[i])
		fmt.Println(g.player.fieldInn[i])
	}
}

func (g *Game) StepComp() bool {
	ok := false
	isGameOver := false

	for i := 0; i < 200; i++ {
		ok, isGameOver = g.player.MakeShot(rand.Intn(10), rand.Intn(10), false)
		if ok {
			g.Print()
			break
		}
	}
	return isGameOver
}

func (g *Game) StepPlayer() bool {
	var hv string
	fmt.Fscan(os.Stdin, &hv)
	i, j := byteToIndex(hv[0], hv[1])
	_, isGameOver := g.comp.MakeShot(i, j, true)
	g.Print()

	return isGameOver
}

func Run() {
	game := CreateGame()
	game.Print()

	isGameOver := false

	for !isGameOver {
		isGameOver = game.StepPlayer()
		if !isGameOver {
			isGameOver = game.StepComp()
		}
	}
	fmt.Println("The Game is Over!!!")
}
