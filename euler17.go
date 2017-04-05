/*
Project Euler Problem #17
https://projecteuler.net/problem=17
*/

package main

import "fmt"

// The smallest length of 1 to 19 (i.e. "one"); Numbers with this length
// are not included in the onetonineteenpositions map
var onetonineteenmin int = 3

// The additional letters for each number in English in addition to onetonineteenmin
// i.e. three has 2 letters in addition to the min (3)
var onetonineteenpositions = map[int]int{
	3:  2,
	4:  1,
	5:  1,
	7:  2,
	8:  2,
	9:  1,
	11: 3,
	12: 3,
	13: 5,
	14: 5,
	15: 4,
	16: 4,
	17: 6,
	18: 5,
	19: 5,
}

// The smallest length of the English representation of numbers 20 to 99
// i.e. "forty" has 5 letters
var secondposmin = 5

// The lengths of English representations in addition to the min.
var secondpositions = map[int]int{
	20: 1,
	30: 1,
	70: 2,
	80: 1,
	90: 1,
}

func onetonineteenlen(num int) int {
	if num < 1 || num > 19 {
		panic("Invalid argument - outside of range 1 to 19")
	}
	return onetonineteenpositions[num] + onetonineteenmin
}

func twentytoninetyninelen(num int) int {
	if num < 20 || num > 99 {
		panic("Invalid argument - outside of range 20 to 99")
	}
	basepos2 := (num / 10) * 10
	rempos2 := num - basepos2
	strtotal := secondpositions[basepos2] + secondposmin
	if rempos2 > 0 {
		strtotal += onetonineteenlen(rempos2)
	}
	return strtotal
}

func onehundredtoonethousandlen(num int) int {
	if num < 100 || num > 1000 {
		panic("Invalid argument - outside of range 100 to 1000")
	}
	if num == 1000 {
		return 11
	}
	basepos3 := (num / 100) * 100
	rempos3 := num - basepos3
	hundredbase := basepos3 / 100
	strtotal := onetonineteenlen(hundredbase) + 7 // i.e "one hundred"
	if rempos3 > 0 && rempos3 < 20 {
		strtotal += onetonineteenlen(rempos3) + 3 // i.e. "and thirteen"
	} else if rempos3 > 19 && rempos3 < 100 {
		strtotal += twentytoninetyninelen(rempos3) + 3 // i.e. and "seventy six"
	}
	return strtotal
}

func main() {
	total := 0
	for i := 1; i <= 1000; i++ {
		switch {
		case i < 20:
			total += onetonineteenlen(i)
		case i >= 20 && i < 100:
			total += twentytoninetyninelen(i)
		case i >= 100:
			total += onehundredtoonethousandlen(i)
		}
	}
	fmt.Println(total)
}
