package interactors

import (
	"github.com/jlarusso/gonads/maybe"
)

type SauceInteractor struct{}

// This particular interactor requires this kind of map as the params
func (si SauceInteractor) Act(params interface{}) maybe.Monad {
	return maybe.Maybe(params).
		Bind(si.Prepare).
		Bind(si.Cook).
		Bind(si.Taste)
}

func (_ SauceInteractor) Prepare(params interface{}, m maybe.Monad) maybe.Monad {
	// NOTE: I'd really like to have some kind of sanitization step where the params coming into
	// the interactor.  The type assertion below would happen there (instead of at the top of each
	// method).  There'd be some kind of context struct specific to this package is passed from
	// function to function.  I couldn't get it working with the monad code because nothing seemed to
	// satisfy interface{}.
	p := params.(map[string]int)

	if p["tomatoes"] < 5 {
		return m.Failure("Not enough tomatoes")
	} else {
		return m.Success(p)
	}
}

func (_ SauceInteractor) Cook(params interface{}, m maybe.Monad) maybe.Monad {
	p := params.(map[string]int)

	if p["heat"] < 80 {
		return m.Failure("Turn up the heat")
	} else {
		return m.Success(p)
	}
}

func (_ SauceInteractor) Taste(params interface{}, m maybe.Monad) maybe.Monad {
	p := params.(map[string]int)

	if p["salt"] < 2 {
		return m.Failure("Thats-a some bland a-sauce")
	} else {
		return m.Success(p)
	}
}
