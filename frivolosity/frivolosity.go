package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var someNumber int
	{
		someNumber = rand.Int()
	}
	fmt.Printf("Hello World! %d\n", someNumber)

	for i := 0; i < 5; i++ {
		fmt.Printf("Normal %d\n", i)
		defer fmt.Printf("Defer %d\n", i)
	}

	t1 := new(SomeType)
	t1.f1 = 3

	fmt.Printf("Foo: %d\n", t1.f1)

	t2 := SomeType{f1: 4}
	fmt.Printf("Foo2: %d\n", t2.f1)

	t2.f1 = 5
	fmt.Printf("Foo3: %d\n", t2.f1)

	//t2 = &SomeType{f1: 6}
	//fmt.Printf("Foo3: %d\n", t2.f1)

	DoSomething(&t2)
	fmt.Printf("Foo4: %d\n", t2.f1)

	DoSomething(&t2)
	fmt.Printf("Foo4: %d\n", t2.f1)

	t2.PointerReceiver()
	fmt.Printf("After PointerReceiver: %s\n", t2)

	t2.ValueReceiver()
	fmt.Printf("After ValueReceiver: %s\n", t2)

	var i1 FooInterface
	i1 = &t2
	fmt.Printf("After interface cast: %s\n", i1)

	i1.PointerReceiver()
	fmt.Printf("After PointerReceiver: %s\n", t2)
}

// Also note the use of multiple variabls
func nakedReturn(a, b int) (multiplyResult, addResult int) {
	multiplyResult = a * b
	addResult = a + b
	return
}

type SomeType struct {
	f1 int
	f2 string
}

func DoSomething(st *SomeType) () {
	st.f1 = 7
	fmt.Printf("DoSomething: %s\n", st)
}

func (st *SomeType) PointerReceiver() () {
	st.f1 = 8
	fmt.Printf("PointerReceiver: %s\n", st)
}

func (st SomeType) ValueReceiver() () {
	st.f1 = 9
	fmt.Printf("ValueReceiver: %s\n", st)
}

type FooInterface interface {
	PointerReceiver()
}
