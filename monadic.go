package main

import "fmt"

type Monad interface {
	Bind(func(interface{}, Monad) Monad) Monad
	Success(interface{}) Monad
	Failure(interface{}) Monad
}

type Maybe struct {
	val *interface{}
}

// implement Bind for Maybe to satisfy Monad
func (m Maybe) Bind(f func(interface{}, Monad) Monad) Monad {
	if m == Failure {
		return Failure
	}
	return f(*m.val, m)
}

// implement Success for Maybe to satisfy Monad
func (m Maybe) Success(a interface{}) Monad {
	return Some(a)
}

func (m Maybe) Failure(a interface{}) Monad {

}

var Failure = Maybe{nil}

func Some(a interface{}) Monad {
	return Maybe{&a}
}

// How to output when doing Println
func (m Maybe) String() string {
	if m == Failure {
		return "Failure"
	}
	return fmt.Sprintf("Some(%v)", *m.val)
}

func main() {
	doubleMe := func(i interface{}, m Monad) Monad {
		return m.Success(2 * i.(int))
	}

	tripleMe := func(i interface{}, m Monad) Monad {
		return m.Success(3 * i.(int))
	}

	oops := func(i interface{}, m Monad) Monad {
		return Failure
	}

	result1 := Some(5).
		Bind(doubleMe).
		Bind(tripleMe)
	fmt.Println(result1)

	result2 := Some(3).
		Bind(oops).
		Bind(doubleMe)
	fmt.Println(result2)

}
