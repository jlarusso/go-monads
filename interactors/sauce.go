package interactors

import (
	"github.com/jlarusso/gonads/maybe"
)

type SauceInteractor struct{}

// This particular interactor requires this kind of map as the params
func (si SauceInteractor) Act(params map[string]int) maybe.Monad {
	return maybe.Maybe(params).
		Bind(si.Sanitize).
		Bind(si.Prepare).
		Bind(si.Cook).
		Bind(si.Taste)
}

func (_ SauceInteractor) Sanitize(params interface{}, m maybe.Monad) maybe.Monad {
	// in order to access it i have to do a type assertion
	return m.Success(params.(map[string]int))
}

func (_ SauceInteractor) Prepare(params map[string]int, m maybe.Monad) maybe.Monad {
	if params["tomatos"] < 5 {
		return m.Failure("Not enough tomatos")
	} else {
		return m.Success(params)
	}
}

func (_ SauceInteractor) Cook(params map[string]int, m maybe.Monad) maybe.Monad {
	if params["heat"] < 80 {
		return m.Failure("Turn up the heat")
	} else {
		return m.Success(params)
	}
}

func (_ SauceInteractor) Taste(params map[string]int, m maybe.Monad) maybe.Monad {
	if params["salt"] < 2 {
		return m.Failure("Thats-a some bland a-sauce")
	} else {
		return m.Success(params)
	}
}
