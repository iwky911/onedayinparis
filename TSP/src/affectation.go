package main

type Stop struct {
	id             int
	next, previous *Stop
}

func bindLeft(a, b *Stop){
	if a!= nil{
		a.next = b
	}
	}
func bindRight(a,b *Stop){
	if(b!=nil) {
		b.previous = a
	}
	}

func swap(a, b *Stop) {
	bindRight( b, a.next)
	bindLeft(a.previous, b)
	bindRight(a, b.next)
	bindLeft(b.previous, a)
	a.next, b.next = b.next, a.next
	a.previous, b.previous = b.previous, a.previous
}

func (s *Stop) start() *Stop{
	for(s.previous != nil) {
		s = s.previous
	}
	return s
	}

func invertRange(a, b *Stop) {
	if a == b {
		return
	}

	before := func(start, end *Stop) bool {
		for start != nil {
			if start == end {
				return true
			}
			start = start.next
		}
		return false
	}(a, b)

	if !before {
		invertRange(b, a)
		return
	}
	left := a.previous
	right := b.next
	bindLeft(left, b)
	bindRight(a, right)
	var current, next *Stop
	current = a
	next = a.next
	for next != nil && current != b {
		nextnext := next.next
		current.previous = next
		next.next = current
		current = next
		next = nextnext
	}
	a.next = right
	b.previous = left
}
