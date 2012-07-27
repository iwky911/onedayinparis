package main

import (
	"fmt"
	"runtime/pprof"
	"flag"
	"os"
)

func (p *Problem) startSolver(nbiter int, initialratio, step float64, out chan Affectation) {
	fmt.Println("launched")
	aff := NewAffectation(303)
	out <- p.SolveWithSimulatedAnnealing(aff, nbiter, initialratio,step)
	}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			println(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	p := NewProblem("matrix.txt", 303)
	c:= make(chan Affectation)
	for i:=0; i<1; i++{
		go p.startSolver(400, 0.6, 0.05, c)
	}
	fmt.Println("hello")
	for i:=0; i<1; i++{
		sol := <-c
		fmt.Println(p.Cost(&sol))
		sol.Dump()
	}
}
