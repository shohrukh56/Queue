package main

import "testing"

func Test_list_add(t *testing.T) {
	q := queue{}
	if q.size != 0 {
		t.Error("new queue length has to be zero, got: ", q.size)
	}
}
func Test_list_add_one(t *testing.T) {
	q := queue{}
	q.equeue("First Person")
	if q.firstElement() != "First Person" {
		t.Error("first in queue is First Person, got:", q.firstElement())
	}
	if q.lastElement() != "First Person" {
		t.Error("last in queue is First Person, got:", q.lastElement())
	}
	if q.size != 1 {
		t.Error("after adding one person size has to be 1, got: ", q.size)
	}
	increment := q.dequeue()
	if increment != "First Person" {
		t.Error("the first person in line is First Person, got: ", increment, )
	}
	if q.size != 0 {
		t.Error("after deleting one element size must be 0, got: ", q.size)
	}
}
