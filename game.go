package main

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	countOneShips = 10
)

type Game struct {
	GameOver bool
	comp   *Battle
	player *Battle
}

func Start() *Game {
	game := Game{GameOver: false, comp: NewBattle(), player: NewBattle()}
	game.comp.Init(countOneShips)
	game.player.Init(countOneShips)
	game.Print()
	return &game
}

func (g *Game) Print() {
	header := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Print(" ")
	fmt.Print(header)
	fmt.Println(header)
	for i := 0; i < len(g.comp.fieldOut); i++ {
		fmt.Print(i)
		if debug {
			fmt.Print(g.comp.fieldInn[i])
			fmt.Println(g.player.fieldInn[i])
		} else {
			fmt.Print(g.comp.fieldOut[i])
			fmt.Println(g.player.fieldInn[i])
		}
	}
}

func getRandIndex(maxIndex int) (int, int) {
	return rand.Intn(maxIndex), rand.Intn(maxIndex)
}

func byteToIndex(h byte, v byte) (int, int) {
	i := v - 48
	j := h - 97
	return int(i), int(j)
}

func (g *Game) MakeShotComp() bool {
	if !g.GameOver {
		for i := 0; i < 200; i++ {
			i, j := getRandIndex(10)
			if g.player.MakeShot(i, j, false) {
				break
			}
		}
		g.GameOver = g.player.GameOver
	}
	return !g.GameOver
}

func (g *Game) MakeShotPlayer() bool {
	if !g.GameOver {
		var yx string
		fmt.Fscan(os.Stdin, &yx)
		if yx == "q" {
			g.GameOver = true
			return false
		} else if yx == "s" {
			g.GameOver = false
			return false
		}
		y, x := byteToIndex(yx[0], yx[1])
		g.comp.MakeShot(y, x, true)
		g.GameOver = g.comp.GameOver
	}
	return !g.GameOver
}

func (g *Game) Complete() {
	g.Print()
	if g.comp.GameOver {
		fmt.Println("The Game is Over!!!\nThe Winner is Player!!!")
	} else if g.player.GameOver {
		fmt.Println("The Game is Over!!!\nThe Winner is Computer!!!")
	}
}

func Run() {
	game := Start()

	for game.MakeShotPlayer() && game.MakeShotComp() {
		game.Print()
	}

	game.Complete()

	if game.GameOver == false {
		Run()
	}
}
