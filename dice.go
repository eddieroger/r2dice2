package main

import (
	"math/rand"
	"time"
)

// Roll rolls a die
func (d *Die) Roll() {

	showing := random(1, d.sides) // do the roll
	d.showing = showing
	d.typeDescription = stringForType(d.dieType)

	r := GetEmptyResult()

	switch d.dieType {
	case DieTypeForce:
		switch showing {
		case 1, 2, 3, 4, 5, 6: // one dark pip
			r.dark = 1
			d.faceDescription = "One Dark Pip"
		case 7, 8: // one light pip
			r.light = 1
			d.faceDescription = "One Light Pip"
		case 9: // two dark pips
			r.dark = 2
			d.faceDescription = "Two Dark Pips"
		case 10, 11, 12: // twp light pips
			r.light = 2
			d.faceDescription = "Two Light Pips"
		}
	}
	d.result = r
}

func random(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func stringForType(dieType DieType) string {
	switch dieType {
	case DieTypeUnknown:
		return "Unknown"
	case DieTypeAbility:
		return "Ability"
	case DieTypeDifficulty:
		return "Difficulty"
	case DieTypeProficieny:
		return "Proficency"
	case DieTypeChallenge:
		return "Challenge"
	case DieTypeBoost:
		return "Boost"
	case DieTypeSetback:
		return "Setback"
	case DieTypeForce:
		return "Force"
	default:
		return "Unknown"
	}
}
