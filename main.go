package main

import "fmt"

func main() {
	forceDie1 := &Die{dieType: DieTypeForce, sides: 12}
	forceDie2 := &Die{dieType: DieTypeForce, sides: 12}
	cup := Cup{}
	cup.AddToCup(forceDie1)
	cup.AddToCup(forceDie2)
	result := cup.Roll()
	fmt.Println("Final Result", result)
}
