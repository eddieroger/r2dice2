package main

// DieType tracks what kind of die we're rolling
type DieType int

const (
	DieTypeUnknown DieType = iota
	DieTypeAbility
	DieTypeDifficulty
	DieTypeProficieny
	DieTypeChallenge
	DieTypeBoost
	DieTypeSetback
	DieTypeForce
)

// Die represents a singular die
type Die struct {
	dieType         DieType
	sides           int
	showing         int
	typeName        string
	faceDescription string
	typeDescription string
	result          *Result
}

// Cup holds and rolls dice. Dice go in a cup
type Cup struct {
	dice []*Die
}

// Result is the result of rolling the dice
type Result struct {
	nothing   int
	advantage int
	success   int
	threat    int
	failure   int
	triumph   int
	despair   int
	light     int
	dark      int
}
