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

func TestPush(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	r1 := Collect(numbers).Push(6, 7, 8).Value()
	r2 := Collect(numbers).Push().Value()
	r3 := Collect(numbers).Push(6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18).Value()

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, r1)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, r2)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, r3)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, numbers)
}

func TestMap(t *testing.T) {
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

func TestFilter(t *testing.T) {
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

func TestMapAndFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	r1 := Collect(numbers).Map(func(item int) int {
		return item + 1
	}).Filter(func(item int) bool {
		return item >= 3
	}).Value()

	assert.Equal(t, []int{3, 4, 5, 6}, r1)
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

	var e2 []Person = []Person{
		{Name: "BZ", Age: 21, Height: 162},
		{Name: "DZ", Age: 30, Height: 199},
		{Name: "EZ", Age: 41, Height: 154},
	}

	r1 := Collect(persons).Map(func(p Person) Person {
		p.Name += "Z"
		p.Age += 1
		p.Height += 2

		return p
	}).Value()

	r2 := Collect(r1).Filter(func(p Person) bool {
		return p.Age > 20
	}).Value()

	assert.Equal(t, e1, r1)
	assert.Equal(t, e2, r2)
}
