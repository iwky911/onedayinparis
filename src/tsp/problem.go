package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"fmt"
)

type Problem struct {
	costMatrix *[][]int
}

func NewProblem(filename string, size int) Problem {
	myfile, _ := ioutil.ReadFile(filename)
	output := make([][]int, 0)
	for i, line := range strings.Split(string(myfile), "\n") {
		if i >= size {
			break
		}
		output = append(output, make([]int, 0))
		for j, nb := range strings.Split(line, " ") {
			if j >= size {
				break
			}
			mynb, _ := strconv.ParseInt(nb, 10, 64)
			output[i] = append(output[i], int(mynb))
		}
	}
	return Problem{&output}
}

func (p *Problem) getCost(a, b int) int {
	return (*p.costMatrix)[a][b]
}

func (p *Problem) Cost(aff *Affectation) int {
	affdata := (*aff).data
	cost := 0
	for current := 0; current < len(affdata)-1; current++ {
		cost += p.getCost(affdata[current], affdata[current+1])
	}
	return cost
}


func (a *Affectation) Dump() {
	fmt.Println(a.data)
}
