package main

import (
	"fmt"
	"testing"

	"github.com/dhendry/kitchen-sink/solitaire/game"
	"github.com/dhendry/kitchen-sink/solitaire/model"
	"github.com/golang/protobuf/jsonpb"
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
	assert.Equal(t, false, v)
	assert.Equal(t, false, ok)

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
	assert.Equal(t, 1, acc(1))
	assert.Equal(t, 2, acc(1))
	assert.Equal(t, 4, acc(2))
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

func TestGameState(t *testing.T) {
	gs := game.NewGameState()
	//gs := &model.GameState{}

	jsonString, ok := (&jsonpb.Marshaler{}).MarshalToString(gs)
	if ok != nil {
		panic(ok)
	}

	fmt.Printf("%v", jsonString)
	//assert.Fail(t, fmt.Sprintf("%+v", gs))
	//assert.Fail(t, jsonString)
}

func TestSlices(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, 10, len(ints))
	assert.Equal(t, 10, cap(ints))

	firstHalf := ints[:5]
	assert.Equal(t, 5, len(firstHalf))
	assert.Equal(t, 10, cap(firstHalf))
	assert.Equal(t, []int{0, 1, 2, 3, 4}, firstHalf)
	assert.True(t, &firstHalf == &firstHalf) // Paranoia

	secondHalf := ints[5:]
	assert.Equal(t, 5, len(secondHalf))
	assert.Equal(t, 5, cap(secondHalf))
	assert.Equal(t, []int{5, 6, 7, 8, 9}, secondHalf)

	//
	// Append to the first half
	//
	firstHalf2 := append(firstHalf, 20)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 20}, firstHalf2)
	assert.Equal(t, 6, len(firstHalf2))
	assert.Equal(t, 10, cap(firstHalf2))

	// This is kinda weird since len (and cap) are value parts of the firstHalf struct
	assert.False(t, &firstHalf == &firstHalf2)
	assert.Equal(t, 5, len(firstHalf))
	assert.Equal(t, 10, cap(firstHalf))

	assert.Equal(t, []int{0, 1, 2, 3, 4, 20, 6, 7, 8, 9}, ints)
	assert.Equal(t, []int{20, 6, 7, 8, 9}, secondHalf)

	secondHalf[1] = 22
	assert.Equal(t, []int{0, 1, 2, 3, 4, 20, 22, 7, 8, 9}, ints)
	assert.Equal(t, []int{20, 22, 7, 8, 9}, secondHalf)

	//
	// Append to the second half
	//
	secondHalf2 := append(secondHalf, 21)
	assert.Equal(t, []int{20, 22, 7, 8, 9, 21}, secondHalf2)

	secondHalf2[2] = 23
	assert.Equal(t, []int{20, 22, 23, 8, 9, 21}, secondHalf2)
	assert.Equal(t, []int{20, 22, 7, 8, 9}, secondHalf) // Still pointing at the original array
	assert.Equal(t, []int{0, 1, 2, 3, 4, 20, 22, 7, 8, 9}, ints)
}
