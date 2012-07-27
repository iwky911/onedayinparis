package main

import(
	"container/heap"
	"fmt"
	"os"
	. "strconv"
	"bufio"
	"strings"
	"io/ioutil"
	)


type Graph struct {
	edges [][]Edge
	}

type Edge struct {
	end int
	cost int
	}

func (g *Graph) Matrix() [][]int {
	sortie := make([][]int, len(g.edges))
	for i,_ := range g.edges {
		sortie[i] = g.Distances(i)
	}
	
	return sortie
}

func (g *Graph) Distances(start int) []int {
	visited := map[int]bool{}
	pq := make(PriorityQueue, 0, len(g.edges)*100)
	pq.AddNode(&Item{start, 0, 0})
	
	distances := make([]int, len(g.edges))
	
	for len(pq)>0 {
		n := heap.Pop(&pq).(*Item)
		if visited[n.value] {
			continue
		}
		visited[n.value] = true
		distances[n.value] = n.priority
		for _, edge := range g.edges[n.value] {
			pq.AddNode(&Item{edge.end, edge.cost + n.priority, 0})
		}
	}
	return distances
}

func saveToFile(filename string, matrix [][]int ) {
	f, err := os.OpenFile(filename, os.O_RDWR| os.O_CREATE | os.O_TRUNC, 666)
	defer f.Close()
	if err != nil {
		println("erreur lors de l'ouverture du fichier")
	}
	writer := bufio.NewWriter(f)
	for _, tab := range matrix{
		for _, j := range tab {
			writer.WriteString(fmt.Sprintf("%d ", j))
		}
		writer.WriteString("\n")
	}
}

func loadGraph(filename string) Graph {
	file,_ := ioutil.ReadFile(filename)
	g := Graph{make([][]Edge, 500)}
	for i,_ := range g.edges {
		println("la")
		g.edges[i] = make([]Edge,0)
	}
	for _,line := range strings.Split(string(file),"\n") {
		//~ println("here")
		if line=="" {
			continue
		}
		var args = strings.Split(line, ",")
		//fmt.Println(args)
		//~ fmt.Println(args[0])
		//~ fmt.Println(args[1])
		//~ fmt.Println(args[2])
		start,_ := ParseInt(args[0], 10, 64)
		end,_ := ParseInt(args[1], 10, 64)
		cost,_ := ParseInt(args[2], 10, 6)
		//fmt.Println(g)
		g.edges[start] = append(g.edges[int(start)], Edge{int(end), int(cost)})
	}
	return g
	
}

func main(){
	g := loadGraph("edges.csv")
	fmt.Println(g.edges)
	saveToFile("matrix.txt", g.Matrix())
	}
