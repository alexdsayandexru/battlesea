package main

import (
    "fmt"
    "os"
)

func printBattleField(c battlefield, m battlefield) {
	header := []string{"a","b","c","d","e","f","g","h","i","j"}
	fmt.Print(" ")
	fmt.Print(header)
	fmt.Println(header)
	for i:=0; i<len(c); i++ {
		fmt.Print(i)
		fmt.Print(c[i])
		fmt.Println(m[i])
	}
}

func byteToIndex(h byte, v byte) (int, int){
	i := v - 48
	j := h - 97
	return int(i), int(j)
}

func main() {
	fmt.Println("")
	battle := CreateBattle()
	battle.Init()
	
	var hv string
	for hv != "end" {
		fmt.Fscan(os.Stdin, &hv)
		i, j := byteToIndex(hv[0],hv[1])
		ok, isGameOver := battle.shot(i, j)
		if ok {
			printBattleField(battle.fieldInn, battle.fieldOut)
		}
		if isGameOver {
			fmt.Println("The Game is Over!!!")
			break
		}
	}
}