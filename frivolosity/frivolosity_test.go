package main

import (
	"fmt"
	"testing"

	"github.com/dhendry/kitchen-sink/solitaire/model"
	"github.com/stretchr/testify/assert"
)

func TestNakedReturn(t *testing.T) {
	fmt.Printf("TestNakedReturn\n")

	mult, add := nakedReturn(5, 6)

	if mult != 30 {
		t.Fail()
	}

	if add != 11 {
		t.Fail()
	}
}

func TestMaps(t *testing.T) {
	fmt.Printf("TestMaps\n")

	m := map[string]bool{
		"java": false,
		"go":   true,
	}

	assert.Equal(t, true, m["go"])
	assert.Equal(t, false, m["java"])
	assert.Equal(t, false, m["fooo"])

	v, ok := m["fooo"]
	assert.Equal(t, false , v)
	assert.Equal(t, false , ok)

}

func getAccumulator() func(int) int {
	var sum int
	return func(i int) int {
		sum += i
		return sum
	}
}

func TestAccumulator(t *testing.T) {
	acc := getAccumulator()
	assert.Equal(t, 1 , acc(1))
	assert.Equal(t, 2 , acc(1))
	assert.Equal(t, 4 , acc(2))
}


func TestNil(t *testing.T) {
	var v1 *SomeType
	assert.True(t, v1.IsNil())

	v1 = &SomeType{}
	assert.False(t, v1.IsNil())

}

func (st *SomeType) IsNil() bool {
	return st == nil
}

func TestSolitaireModel(t *testing.T) {
	assert.Equal(t, 2, int(model.Suit_DIAMONDS))
}
