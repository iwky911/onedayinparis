package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	. "strconv"
	"strings"
)

type Graph struct {
	Edges [][]Edge
}

type Edge struct {
	end  int
	cost int
	line string
}

func (g *Graph) Matrix() [][]int {
	sortie := make([][]int, len(g.Edges))
	for i, _ := range g.Edges {
		sortie[i] = g.Distances(i)
	}

	return sortie
}

func (g *Graph) Distances(start int) []int {
	visited := map[int]bool{}
	pq := make(PriorityQueue, 0, len(g.Edges)*100)
	pq.AddNode(&Item{value: start, priority: 0, index: 0})

	distances := make([]int, len(g.Edges))

	for len(pq) > 0 {
		n := heap.Pop(&pq).(*Item)
		if visited[n.value] {
			continue
		}
		visited[n.value] = true
		distances[n.value] = n.priority
		for _, edge := range g.Edges[n.value] {
			pq.AddNode(&Item{value: edge.end, priority: edge.cost + n.priority, index: 0})
		}
	}
	return distances
}

func SaveMatrixToFile(filename string, matrix [][]int) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 666)
	defer f.Close()
	if err != nil {
		println("erreur lors de l'ouverture du fichier")
	}
	writer := bufio.NewWriter(f)
	for _, tab := range matrix {
		for _, j := range tab {
			writer.WriteString(fmt.Sprintf("%d ", j))
		}
		writer.WriteString("\n")
	}
}

func LoadGraph(filename string) Graph {
	file, _ := ioutil.ReadFile(filename)
	g := Graph{make([][]Edge, 500)}
	for i, _ := range g.Edges {
		// println("la")
		g.Edges[i] = make([]Edge, 0)
	}
	for _, line := range strings.Split(string(file), "\n") {
		//~ println("here")
		if line == "" {
			continue
		}
		var args = strings.Split(line, ",")
		start, _ := Atoi(args[0])
		end, _ := Atoi(args[1])
		cost, _ := Atoi(args[2])
		line := args[3]
		//fmt.Println(g)
		g.Edges[start] = append(g.Edges[int(start)], Edge{int(end), int(cost), line})
	}
	return g
}

func ParseNodePath(filename string) []int {
	file, _ := ioutil.ReadFile(filename)
	strNodes := strings.Split(string(file[1:len(file)-1]), " ")
	ids := make([]int, len(strNodes))
	for i, n := range strNodes {
		nI, _ := Atoi(n)
		ids[i] = nI
	}
	return ids
}

func (g *Graph) backtrack(v *Item) []int {
	a := make([]int, 0)
	i := 0
	for c := v; c.parent != nil; c = c.parent {
		a = append(a, c.parent.value)
		// a = append(a[:i], append([]int{c.parent.value}, a[i:]...)...)
		i += 1
	}
	b := make([]int, len(a))
	for i := 0; i < len(b); i++ {
		b[len(b)-1-i] = a[i]
	}
	// fmt.Println("Path:", b)
	return b
}

func (g *Graph) ShortestPath(start int, end int) []int {
	visited := map[int]bool{}
	pq := make(PriorityQueue, 0, len(g.Edges)*100)
	pq.AddNode(&Item{start, 0, 0, nil})
	distances := make([]int, len(g.Edges))

	for len(pq) > 0 {
		n := heap.Pop(&pq).(*Item)
		if visited[n.value] {
			continue
		}
		if n.value == end {
			// fmt.Println("done!")
			return g.backtrack(n)
		}
		visited[n.value] = true
		distances[n.value] = n.priority
		for _, edge := range g.Edges[n.value] {
			pq.AddNode(&Item{edge.end, edge.cost + n.priority, 0, n})
		}
	}
	log.Fatal("No path between", start, "and", end, "found")
	return []int{}
}

func (g *Graph) FullPath(nodeIds []int) []int {
	fullIds := make([]int, 0)
	for i := 0; i < len(nodeIds)-1; i++ {
		fullIds = append(fullIds, g.ShortestPath(nodeIds[i], nodeIds[i+1])...)
	}
	fullIds = append(fullIds, nodeIds[len(nodeIds)-1])
	return fullIds
}

func SavePathToFile(filename string, nodes []int) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 666)
	defer f.Close()
	if err != nil {
		println("erreur lors de l'ouverture du fichier")
	}
	writer := bufio.NewWriter(f)
	for _, j := range nodes {
		writer.WriteString(fmt.Sprintf("%d ", j))
	}
	writer.Flush()
}
