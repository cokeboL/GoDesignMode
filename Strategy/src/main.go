package main

import (
	"Duck"
	"fmt"
)

func DuckAction(duck Duck.Duck) {
	fmt.Println("\n----------------------------")
	duck.Description()
	duck.Fly()
	duck.Sing()
	fmt.Println("----------------------------")
}

func main() {
	gDuck := &Duck.GreenDuck{}
	wDuck := &Duck.WildDuck{}

	DuckAction(gDuck)

	DuckAction(wDuck)
}
