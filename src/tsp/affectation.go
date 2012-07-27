package main

import (
	"fmt"
	"math/rand"
	//~ "reflect"
	"hash/crc32"
	//~ "os"
)

type AffectData []int

type Affectation struct {
	data AffectData
	name string
}

func (a *Affectation) Correct() bool {
	m := make(map[int]bool)
	for _, v := range a.data {
		if m[v] {
			//println(v)
		}
		m[v] = true
	}
	return len(a.data) == len(m)
}

func (a *Affectation) copy() Affectation {
	data := make([]int, len(a.data))
	copy(data, a.data)
	return Affectation{AffectData(data), a.name}
}

func (a *Affectation) FindMissing() {
	m := make(map[int]bool)
	for _, v := range a.data {
		m[v] = true
	}
	for i := range a.data {
		if !m[i] {
			println(i)
		}
	}
	//return len(a.data)==len(m)
}

func (a *Affectation) imprint() uint32 {
	return crc32.ChecksumIEEE([]byte(fmt.Sprint(a)))
}

func NewAffectation(size int) Affectation {
	aff := make([]int, size)
	for i, _ := range aff {
		aff[i] = i
	}
	return Affectation{aff, ""}
}

func RandomAffectation(size int) Affectation {
	return Affectation{rand.Perm(size), ""}
}

func (a *Affectation) getPossibleAffectation(possibility *[]Affectation) {
	k := 0
	for i := 0; i < len((*a).data); i++ {
		if i > 0 {
			(*a).cutAt(i, &(*possibility)[k])
			k = k + 1
		}
		for j := i + 1; j < len((*a).data); j++ {
			(*a).swap(i, j, &(*possibility)[k])
			k = k + 1
			(*a).invertRange(i, j, &(*possibility)[k])
			k = k + 1
		}
		//~ if !(*possibility)[k-1].Correct() {
		//~ fmt.Println("incorrect!!! ",k, (*possibility)[k].name, a.data)
		//~ fmt.Println((*possibility)[k].data)
		//~ os.Exit(6)
		//~ }
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
}

func (p *Problem) getCheapest(possibilities *[]Affectation) *Affectation {
	cost := 1000 * 1000 * 1000
	var record Affectation
	for _, aff := range *possibilities {
		newcost := p.Cost(&aff)
		//~ if !aff.Correct() {
		//~ fmt.Println("incorrect!!!", aff.name)
		//~ fmt.Println(aff.data)
		//~ os.Exit(6)
		//~ }
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
	//~ mincost := p.Cost(startpoint)
	//~ 
	//~ for i := 0; i < 4; i++ {
	//~ newstop := p.getCheapest(startpoint.getPossibleAffectation())
	//~ if mincost > p.Cost(newstop) {
	//~ mincost = p.Cost(newstop)
	//~ bestsolution = newstop
	//~ }
	//~ }
	return bestsolution
}

func (a *Affectation) hop() Affectation {
	n := len(a.data)
	k := rand.Intn(2)
	i := rand.Intn(n - 1)
	j := i + 1 + rand.Intn(n-i-1)
	sortie := Affectation{AffectData(make([]int, n)), ""}
	switch {
	case i > 0 && k == 0:
		a.cutAt(i, &sortie)
	case k == 1 || k == 0 && i == 0:
		a.swap(i, j, &sortie)
	case k == 2:
		a.invertRange(i, j, &sortie)
	}
	return sortie
}

func (p *Problem) SolveWithSimulatedAnnealing(startpoint Affectation, nbiteration int, initialproba, rate float64) Affectation {
	bestsolution := startpoint
	mincost := p.Cost(&bestsolution)

	var previous string
	var proba = initialproba
	possibleSolutions := make([]Affectation, len(startpoint.data)*(len(startpoint.data))-1)
	for i := range possibleSolutions {
		possibleSolutions[i] = Affectation{AffectData(make([]int, len(startpoint.data))), ""}
	}
	currentSolution := startpoint
	for i := 0; i < nbiteration; i++ {
		previous = currentSolution.name
		currentSolution.getPossibleAffectation(&possibleSolutions)

		if rand.Float64() < proba {
			currentSolution = p.getCheapest(&possibleSolutions).copy()
		} else {
			currentSolution = possibleSolutions[rand.Intn(len(possibleSolutions))].copy()
		}

		if p.Cost(&currentSolution) < mincost {
			bestsolution = currentSolution
			mincost = p.Cost(&currentSolution)
			println("current cost:", p.Cost(&currentSolution), "with", currentSolution.name)
		}
		if currentSolution.name == previous {
			currentSolution = currentSolution.hop()
			proba = 0.5
			//currentSolution = bestsolution
			fmt.Println("proba: ", proba, "\t")
		}

		proba = proba + rate
		if proba > 1 {
			proba = 1
		}
	}
	return bestsolution
}
