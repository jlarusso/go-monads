package main

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
	return m // don't run the current function in the chain, just pass it through.
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

func main() {
	doubleMe := func(i interface{}, m Monad) Monad {
		return m.Success(2 * i.(int))
	}

	tripleMe := func(i interface{}, m Monad) Monad {
		return m.Success(3 * i.(int))
	}

	oops := func(i interface{}, m Monad) Monad {
		return m.Failure("oooooops")
	}

	result1 := Maybe(5).
		Bind(doubleMe).
		Bind(tripleMe)
	fmt.Println(result1)
	// => Success(30)

	result2 := Maybe(3).
		Bind(oops).
		Bind(doubleMe)
	fmt.Println(result2)
	// => Failure(oooooops)

}
