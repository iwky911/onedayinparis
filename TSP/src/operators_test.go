package main

import (
	"fmt"
	"reflect"
	"testing"
)
//~ 
//~ func TestSwap(t *testing.T) {
	//~ a := Affectation([]int{0, 1, 2, 3})
	//~ a = a.swap(0, 2)
//~ 
	//~ if !reflect.DeepEqual(a, Affectation([]int{2, 1, 0, 3})) {
		//~ fmt.Println(a)
		//~ fmt.Println(reflect.DeepEqual(a, []int{2, 1, 0, 3}))
		//~ t.Fail()
	//~ }
//~ }

func TestRevert(t *testing.T) {
	a := Affectation([]int{0, 1, 2, 3})
	a = a.invertRange(0, 3)
	if !reflect.DeepEqual(a, Affectation([]int{3, 2, 1, 0})) {
		t.Fail()
	}
}

func TestCostComputationGiveCostSum(t *testing.T) {
	a := Affectation([]int{0, 1, 2})
	p := Problem{&[][]int{{0, 1, 2}, {2, 0, 2}, {1, 1, 0}}}
	if p.Cost(&a) != 3 {
		t.Fail()
	}
}

func TestChooseSmallerCost(t *testing.T){
	s := Affectation([]int{0,1,2})
	p := Problem{&[][]int{{0, 1, 2}, {2,0,2}, {1,1,0} }}
	
	sol := p.solve(&s)
	if !reflect.DeepEqual(sol, Affectation([]int{2, 0, 1})) {
		fmt.Println(sol)
		t.Fail()	
	}
}

func TestOperatorDoNotChangeFirstAffectation(t *testing.T){
	var s = Affectation([]int{0,1,2})
	b := s.swap(1,2)
	if s[1]==2 && b[1]==2{
		t.Errorf("swap modified s")
	}
	}

func TestCutWorks(t *testing.T){
	a := Affectation([]int{1,2,3})
	b := a.cutAt(1)
	if !reflect.DeepEqual(b, Affectation([]int{2,3,1})) {
		fmt.Println(a, b)
		t.Errorf("not equals! ")
	}
	}
