package main

import(
	//~ "fmt"
	
)

type Affectation []int

type Problem struct {
  costMatrix *[][]int
  
}

func (p *Problem) getCost(a, b int) int {
	return (*p.costMatrix)[a][b]
}

func (p *Problem) Cost(a *Affectation) int {
	cost := 0
	for current := 0; current < len(*a)-1; current++ {
		cost += p.getCost((*a)[current], (*a)[current+1])
	}
	return cost
}

func (a *Affectation) getPossibleAffectation() *[]Affectation {
	possibility := make([]Affectation, 0)
	for i:=0; i<len(*a); i++ {
		if i>0 {
			possibility = append(possibility, (*a).cutAt(i))
		}
		for j:= i+1; j<len(*a); j++ {
			possibility = append(possibility, (*a).swap(i,j))
			possibility = append(possibility, (*a).invertRange(i,j))
		}
	}
	return &possibility
}

func (p *Problem) getCheapest(possibilities *[]Affectation) *Affectation {
	cost := 1000 * 1000 * 1000
	var record *Affectation
	for _, aff := range *possibilities {
		if p.Cost(&aff) < cost {
			record = &aff
		}
	}
	return record
}

func (p *Problem) solve(startpoint *Affectation) *Affectation {
	bestsolution := startpoint
	mincost := p.Cost(startpoint)

	for i := 0; i < 4; i++ {
		newstop := p.getCheapest(startpoint.getPossibleAffectation())
		if mincost > p.Cost(newstop) {
			mincost = p.Cost(newstop)
			bestsolution = newstop
		}
	}
	return bestsolution
}
