package main

import (
	"fmt"
	"github.com/jlarusso/gonads/interactors"
)

func main() {
	params := make(map[string]int)
	params["tomatos"] = 1
	params["heat"] = 100
	params["salt"] = 2

	result1 := interactors.SauceInteractor{}.Act(params)
	fmt.Println(result1) // => Failure(Not enough tomatos)

	params := make(map[string]int)
	params["tomatos"] = 10
	params["heat"] = 20
	params["salt"] = 25

	result1 := interactors.SauceInteractor{}.Act(params)
	fmt.Println(result1) // => Failure(Turn up the heat)

	params := make(map[string]int)
	params["tomatos"] = 5
	params["heat"] = 100
	params["salt"] = 2

	result1 := interactors.SauceInteractor{}.Act(params)
	fmt.Println(result1) // => Success(30)
}
