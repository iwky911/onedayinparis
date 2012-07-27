package main

import (
	"fmt"
	"math/rand"
	//~ "reflect"
	"hash/crc32"
)

type AffectData []int

type Affectation struct{
	data AffectData
	name string
	}

func (a *Affectation ) imprint() uint32 {
	return crc32.ChecksumIEEE([]byte(fmt.Sprint(a)))
}

func NewAffectation(size int) Affectation {
	aff := make([]int, size)
	for i, _ := range aff {
		aff[i] = i
	}
	return Affectation{aff,""}
}

func RandomAffectation(size int) Affectation {
	return Affectation{rand.Perm(size),""}
}

func (a *Affectation) getPossibleAffectation() *[]Affectation {
	possibility := make([]Affectation, 0)
	for i := 0; i < len((*a).data); i++ {
		if i > 0 {
			possibility = append(possibility, (*a).cutAt(i))
		}
		for j := i + 1; j < len((*a).data); j++ {
			possibility = append(possibility, (*a).swap(i, j))
			possibility = append(possibility, (*a).invertRange(i, j))
		}
	}
	//~ aff := possibility[1]
	//~ for i, _ := range possibility[:samplesize] {
		//~ n := rand.Intn(len(possibility)-i)
		//~ //fmt.Println(possibility[i].imprint(), possibility[i+n].imprint())
		//~ possibility[i+n], possibility[i] = possibility[i], possibility[i+n]
		//~ //fmt.Println(possibility[i].imprint(),possibility[i+n].imprint())
	//~ }
	//~ out := possibility//[:samplesize]
	//~ return &out
	return &possibility
}

func (p *Problem) getCheapest(possibilities *[]Affectation) *Affectation {
	cost := 1000 * 1000 * 1000
	var record Affectation
	for _, aff := range *possibilities {
		newcost := p.Cost(&aff)
		if newcost < cost {
			record = aff
			//~ println("found:", aff.name, newcost, cost)
			cost = newcost
		}
	}
	
	//println("return:", record.name, cost)
	return &record
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

func (p *Problem) SolveWithSimulatedAnnealing(startpoint *Affectation, nbiteration int, initialproba, rate float64) *Affectation {
	bestsolution := startpoint
	mincost := p.Cost(bestsolution)
	
	var previous string
	var proba = initialproba
	
	currentSolution := startpoint
	for i := 0; i < nbiteration; i++ {
		//println("iter nb: ", i, " proba: ",proba)
		previous = currentSolution.name
		possibleSolutions := currentSolution.getPossibleAffectation()
		if rand.Float64() < proba {
			currentSolution = p.getCheapest(possibleSolutions)
		} else {
			currentSolution = &(*possibleSolutions)[rand.Intn(len(*possibleSolutions))]
		}

		//println("current cost:",p.Cost(currentSolution), "with", currentSolution.name)
		//~ println(currentSolution.name)
		//~ tmp := currentSolution.cutAt(1)
		//~ println("possible les cost ?", p.Cost(&tmp))
		if p.Cost(currentSolution) < mincost {
			bestsolution = currentSolution
			mincost = p.Cost(currentSolution)
		}
		if currentSolution.name == previous {
			//fmt.Println("resetting the proba: stuck on minimum")
			proba = 00
			//currentSolution = bestsolution
			fmt.Println("proba: ",proba)
		}
		if i%100==0 {
			fmt.Println("iter: ",i)
		}
		proba = proba + rate
		if proba > 1 {
			proba = 1
		}
	}
	return bestsolution
}
