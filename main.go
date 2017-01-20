package main

import (
	"fmt"
	"github.com/jlarusso/gonads/interactors"
)

func main() {
	p1 := make(map[string]int)
	p1["tomatoes"] = 1
	p1["heat"] = 100
	p1["salt"] = 2

	result1 := interactors.SauceInteractor{}.Act(p1)
	fmt.Println(result1) // => Failure(Not enough tomatoes)

	p2 := make(map[string]int)
	p2["tomatoes"] = 10
	p2["heat"] = 20
	p2["salt"] = 25

	result2 := interactors.SauceInteractor{}.Act(p2)
	fmt.Println(result2) // => Failure(Turn up the heat)

	p3 := make(map[string]int)
	p3["tomatoes"] = 5
	p3["heat"] = 100
	p3["salt"] = 2

	result3 := interactors.SauceInteractor{}.Act(p3)
	fmt.Println(result3) // => Success(30)
}
