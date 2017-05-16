package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {

	l := &LinkedList{}

	l.Append("Oompa")
	assert.Equal(t, l.Length(), 1)
	assert.Equal(t, l.head.data, "Oompa")
	assert.Equal(t, l.last.data, "Oompa")

	l.Append("Loompa")
	assert.Equal(t, l.Length(), 2)
	assert.Equal(t, l.head.data, "Oompa")
	assert.Equal(t, l.last.data, "Loompa")

	l.Append("I")
	l.Append("Love")
	l.Append("Natel")
	assert.Equal(t, l.Length(), 5)
	assert.Equal(t, l.head.data, "Oompa")
	assert.Equal(t, l.last.data, "Natel")

	assert.Equal(t, l.Search("Oompa"), true)
	assert.Equal(t, l.Search("peidei"), false)
}

func TestPreppend(t *testing.T) {
	l := &LinkedList{}

	l.Preppend("Oompa")
	assert.Equal(t, l.Length(), 1)
	assert.Equal(t, l.head.data, "Oompa")
	assert.Equal(t, l.last.data, "Oompa")

	l.Preppend("Loompa")
	assert.Equal(t, l.Length(), 2)
	assert.Equal(t, l.head.data, "Loompa")
	assert.Equal(t, l.last.data, "Oompa")

	l.Preppend("biguebigole")
	assert.Equal(t, l.head.data, "biguebigole")
	assert.Equal(t, l.last.data, "Oompa")
	assert.Equal(t, l.Length(), 3)

	assert.Equal(t, l.Search("biguebigole"), true)
	assert.Equal(t, l.Search("pacoca"), false)
}

func TestInsert(t *testing.T) {
	l := &LinkedList{}

	l.Insert("I wanna go to the pub", 3)
	assert.Equal(t, l.head.data, "I wanna go to the pub")
	assert.Equal(t, l.last.data, "I wanna go to the pub")
	assert.Equal(t, l.Length(), 1)
	assert.Equal(t, l.Search("I wanna go to the pub"), true)

	l.Insert("Say My Name", -1)
	assert.Equal(t, l.head.data, "Say My Name")
	assert.Equal(t, l.last.data, "I wanna go to the pub")
	assert.Equal(t, l.Length(), 2)

	l.Insert("middle", 1)
	assert.Equal(t, l.head.data, "Say My Name")
	assert.Equal(t, l.last.data, "I wanna go to the pub")
	assert.Equal(t, l.Search("middle"), true)
	assert.Equal(t, l.Length(), 3)
}

func TestSearch(t *testing.T) {
	l := &LinkedList{}

	assert.Equal(t, l.Search("data"), false)
	l.Insert("data", 0)
	assert.Equal(t, l.Search("data"), true)
}

func TestLength(t *testing.T) {
	l := &LinkedList{}

	assert.Equal(t, l.Length(), 0)

	l.Append("one")
	l.Append("two")
	l.Append("three")
	assert.Equal(t, l.Length(), 3)
}

func TestRemove(t *testing.T) {
	l := &LinkedList{}

	l.Append("one")
	l.Append("two")
	l.Append("three")
	l.Append("four")
	l.Append("five")

	assert.Equal(t, l.Search("one"), true)
	l.Remove("one")
	assert.Equal(t, l.Search("one"), false)

	assert.Equal(t, l.Search("five"), true)
	l.Remove("five")
	assert.Equal(t, l.Search("five"), false)

	assert.Equal(t, l.Search("three"), true)
	l.Remove("three")
	assert.Equal(t, l.Search("three"), false)
}

func TestPop(t *testing.T) {
	l := &LinkedList{}

	l.Append("one")
	l.Append("two")
	l.Append("three")
	l.Append("four")
	l.Append("five")

	assert.Equal(t, l.Search("one"), true)
	l.Pop(0)
	assert.Equal(t, l.Search("one"), false)

	assert.Equal(t, l.Search("five"), true)
	l.Pop(3)
	assert.Equal(t, l.Search("five"), false)

	assert.Equal(t, l.Search("three"), true)
	l.Pop(1)
	assert.Equal(t, l.Search("three"), false)

}

func TestIsEmpty(t *testing.T) {
	l := &LinkedList{}

	assert.Equal(t, l.IsEmpty(), true)

	l.Append("one")
	assert.Equal(t, l.IsEmpty(), false)

}
