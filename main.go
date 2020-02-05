package main

func main() {
}
type Queue interface {
	len() int
	first() *element
	last() *element
	equeue(value interface{})
	dequeue() interface{}
}
type queue struct {
	first *element
	last  *element
	size  int
}
type element struct {
	previous *element
	next     *element
	value    interface{}
}

func (receiver *queue) firstElement() interface{} {
	return receiver.first.value
}
func (receiver *queue) lastElement() interface{} {
	return receiver.last.value
}
func (receiver *queue) dequeue() interface{} {
	if receiver.size == 0 {
		return nil
	}
	if receiver.size == 1 {
		current :=receiver.first.value
		receiver.first = nil
		receiver.last = nil
		receiver.size--
		return current
	}
	current :=receiver.first.value
	receiver.first = receiver.first.next
	receiver.first.previous = nil
	receiver.size--
	return current
}
func (receiver *queue) equeue(value interface{}) {
	if receiver.size == 0 {
		current := &element{
			next:     nil,
			value:    value,
			previous: nil,
		}
		receiver.last = current
		receiver.first = current
		receiver.size++
		return
	}
	previous:=receiver.last
	current :=&element{
		previous: previous,
		next:     nil,
		value:    value,
	}
	previous.next=current
	receiver.last=current
	receiver.size++
}

