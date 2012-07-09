package main

import (
	//"fmt"
	"testing"
	"reflect"
)

func createStopsFromList(l []int) *Stop {
	sortie := make([]Stop, len(l))
	for i, v := range l {
		sortie[i] = Stop{v, nil, nil}
	}
	for i := 0; i < len(l)-1; i++ {
		sortie[i].next = &sortie[i+1]
		sortie[i+1].previous = &sortie[i]
	}
	return &sortie[0]
}

func toInts(s *Stop) []int {
	sortie := make([]int, 0)
	for current := s; current != nil; current = current.next {
		sortie = append(sortie, current.id)
	}
	println(len(sortie))
	return sortie
}

func TestSwap(t *testing.T) {
	s := createStopsFromList([]int{0, 1, 2, 3})
	swap(s, s.next.next)
	
	if !reflect.DeepEqual(toInts(s.start()),[]int{2,1,0,3}) {
		t.Fail()
	}
}

func TestRevert(t *testing.T) {
		s := createStopsFromList([]int{0, 1, 2, 3})
		invertRange(s, s.next.next.next)
		if !reflect.DeepEqual(toInts(s.start()),[]int{3,2,1,0}) {
			t.Fail()
		}
	}
