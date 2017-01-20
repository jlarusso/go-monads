package monads

import "fmt"

type Maybe interface {
	Bind(func(interface{}, Maybe) Maybe) Maybe
	Success(interface{}) Maybe
	Failure(interface{}) Maybe
}

// https://golang.org/doc/effective_go.html#embedding
type Just struct {
	value *interface{}
}

type Success struct {
	Just
}

type Failure struct {
	Just
}

func (m Success) Bind(f func(interface{}, Maybe) Maybe) Maybe {
	return f(*m.value, m) // execute the function and return the monad that IT returns
}

func (m Failure) Bind(f func(interface{}, Maybe) Maybe) Maybe {
	return m // don't run the current function in the chain, just keep going
}

func (m Just) Success(a interface{}) Maybe {
	return Success{Just{&a}}
}

func (m Just) Failure(a interface{}) Maybe {
	return Failure{Just{&a}}
}

func Some(a interface{}) Maybe {
	return Success{Just{&a}}
}

// How to output when doing Println
func (m Success) String() string {
	return fmt.Sprintf("Success(%v)", *m.value)
}

// How to output when doing Println
func (m Failure) String() string {
	return fmt.Sprintf("Failure(%v)", *m.value)
}
