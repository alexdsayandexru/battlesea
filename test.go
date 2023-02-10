package main

import (
	"fmt"
	"math/rand"
	"os"
)

func printBattleField(c battlefield, m battlefield) {
	header := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Print(" ")
	fmt.Print(header)
	fmt.Println(header)
	for i := 0; i < len(c); i++ {
		fmt.Print(i)
		fmt.Print(c[i])
		fmt.Println(m[i])
	}
}

func Test() {
	battle := CreateBattle()
	battle.Init()

	var hv string
	for hv != "end" {
		fmt.Fscan(os.Stdin, &hv)
		i, j := byteToIndex(hv[0], hv[1])
		ok, isGameOver := battle.MakeShot(i, j, true)
		if ok {
			printBattleField(battle.fieldInn, battle.fieldOut)
		}
		if isGameOver {
			fmt.Println("The Game is Over!!!")
			break
		}
	}
}

func Test2() {
	battle := CreateBattle()
	battle.Init()

	isGameOver := false
	for s := 0; s < 200; s++ {
		_, isGameOver = battle.MakeShot(rand.Intn(10), rand.Intn(10), false)

		if isGameOver {
			break
		}
	}
	printBattleField(battle.fieldInn, battle.fieldOut)

	if isGameOver {
		fmt.Println("The Game is Over!!!")
	}
}
