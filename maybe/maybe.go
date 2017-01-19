package maybe

import "fmt"

type Monad interface {
	Bind(func(interface{}, Monad) Monad) Monad
	Success(interface{}) Monad
	Failure(interface{}) Monad
}

// https://golang.org/doc/effective_go.html#embedding
type Return struct {
	value *interface{}
}

type Success struct {
	Return
}

type Failure struct {
	Return
}

func (m Success) Bind(f func(interface{}, Monad) Monad) Monad {
	return f(*m.value, m) // execute the function and return the monad that IT returns
}

func (m Failure) Bind(f func(interface{}, Monad) Monad) Monad {
	return m // don't run the current function in the chain, just keep going
}

func (m Return) Success(a interface{}) Monad {
	return Success{Return{&a}}
}

func (m Return) Failure(a interface{}) Monad {
	return Failure{Return{&a}}
}

func Maybe(a interface{}) Monad {
	return Success{Return{&a}}
}

// How to output when doing Println
func (m Success) String() string {
	return fmt.Sprintf("Success(%v)", *m.value)
}

// How to output when doing Println
func (m Failure) String() string {
	return fmt.Sprintf("Failure(%v)", *m.value)
}
