package main

func GetEmptyResult() *Result {
	// zero out return values
	nothing := 0
	advantage := 0
	success := 0
	threat := 0
	failure := 0
	triumph := 0
	despair := 0
	light := 0
	dark := 0

	return &Result{
		nothing,
		advantage,
		success,
		threat,
		failure,
		triumph,
		despair,
		light,
		dark,
	}
}

func AddResults(r1 *Result, r2 *Result) *Result {
	r := GetEmptyResult()
	r.nothing = r1.nothing + r2.nothing
	r.advantage = r1.advantage + r2.advantage
	r.success = r1.success + r2.success
	r.threat = r1.threat + r2.threat
	r.failure = r1.failure + r2.failure
	r.triumph = r1.triumph + r2.triumph
	r.despair = r1.despair + r2.despair
	r.light = r1.light + r2.light
	r.dark = r1.dark + r2.dark

	return r
}
