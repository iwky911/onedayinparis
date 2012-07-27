package main

import (
	"fmt"
)

func (p *Problem) startSolver(nbiter int, initialratio, step float64, out chan *Affectation) {
	fmt.Println("launched")
	aff := RandomAffectation(303)
	out <- p.SolveWithSimulatedAnnealing(&aff, nbiter, initialratio,step)
	}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	p := NewProblem("matrix.txt", 303)
	c:= make(chan *Affectation)
	for i:=0; i<3; i++{
		go p.startSolver(400, 0.6, 0.05, c)
	}
	fmt.Println("hello")
	for a:= range c {
		fmt.Println(p.Cost(a))
	}
}
