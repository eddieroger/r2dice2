package main

import "fmt"

// AddToCup adds a die to the cup
func (c *Cup) AddToCup(die *Die) {
	c.dice = append(c.dice, die)
}

// Roll rolls the cup, which is really rolling the dice and resolving
func (c *Cup) Roll() *Result {
	finalResult := GetEmptyResult()
	for _, die := range c.dice {

		die.Roll()
		output := fmt.Sprintf("%s: showing %s", die.typeDescription, die.faceDescription)
		fmt.Println(output)
		finalResult = AddResults(finalResult, die.result)
	}
	return finalResult
}
