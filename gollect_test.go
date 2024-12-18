package gollect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name   string
	Height int
	Age    int
}

func TestMapSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	r1 := Collect(numbers).Map(func(value int) int {
		return value + 1
	}).Value()
	r2 := Collect(numbers).Map(func(value int) int {
		return value * 2
	}).Value()

	assert.Equal(t, []int{2, 3, 4, 5, 6}, r1)
	assert.Equal(t, []int{2, 4, 6, 8, 10}, r2)
}

func TestFilterSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	r1 := Collect(numbers).Filter(func(value int) bool {
		return value >= 3
	}).Value()
	r2 := Collect(numbers).Filter(func(value int) bool {
		return value < 3
	}).Value()

	assert.Equal(t, []int{3, 4, 5}, r1)
	assert.Equal(t, []int{1, 2}, r2)
}

func TestMapAndFilterSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	x1 := Collect(numbers).Map(func(item int) int {
		return item + 1
	}).Filter(func(item int) bool {
		return item >= 3
	}).Value()

	assert.Equal(t, []int{3, 4, 5, 6}, x1)
}

func TestInteractWithStruct(t *testing.T) {
	var persons []Person = []Person{
		{Name: "A", Age: 15, Height: 172},
		{Name: "B", Age: 20, Height: 160},
		{Name: "C", Age: 8, Height: 180},
		{Name: "D", Age: 29, Height: 197},
		{Name: "E", Age: 40, Height: 152},
	}

	var e1 []Person = []Person{
		{Name: "AZ", Age: 16, Height: 174},
		{Name: "BZ", Age: 21, Height: 162},
		{Name: "CZ", Age: 9, Height: 182},
		{Name: "DZ", Age: 30, Height: 199},
		{Name: "EZ", Age: 41, Height: 154},
	}

	r1 := Collect(persons).Map(func(p Person) Person {
		p.Name += "Z"
		p.Age += 1
		p.Height += 2

		return p
	}).Value()

	assert.Equal(t, e1, r1)
}
