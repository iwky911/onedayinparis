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
	Edges      [][]Edge
	switchCost int
	nodeMax    int
}

type Edge struct {
	end   int
	cost  int
	lines []string
}

func (g *Graph) Matrix() [][]int {
	sortie := make([][]int, len(g.Edges))
	for i := 0; i < g.nodeMax; i++ {
		sortie[i] = g.Distances(i, g.nodeMax)
	}

	return sortie
}

func (g *Graph) Distances(start, maxId int) []int {
	visited := map[int]bool{}
	pq := make(PriorityQueue, 0) //, len(g.Edges)*100
	// pq.AddNode(&Item{value: start, priority: 0, index: 0, lines: nil})
	pq.AddNode(&Item{value: start, priority: 0, index: 0})
	distances := make([]int, len(g.Edges))
	count := 0
	for len(pq) > 0 {
		n := heap.Pop(&pq).(*Item)
		if visited[n.value] {
			continue
		}
		count++
		if count > maxId {
			break
		}
		visited[n.value] = true
		distances[n.value] = n.priority
		for _, edge := range g.Edges[n.value] {
			switchCost := 0
			// switchCost, linesTraveled := g.calcSwitchCost(n, n.parent)
			// _, newlines := match(edge.lines, linesTraveled)
			// pq.AddNode(&Item{edge.end, switchCost + edge.cost + n.priority, 0, newlines, n})
			pq.AddNode(&Item{edge.end, switchCost + edge.cost + n.priority, 0, n})
			// pq.AddNode(&Item{value: edge.end, priority: switchCost + edge.cost + n.priority, index: 0, lines: edge.lines, parent: n})
		}
	}
	println("count", count)
	return distances
}

func SaveMatrixToFile(filename string, matrix [][]int) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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
		writer.Flush()
	}
}

func LoadGraph(filename string, switchCost int) Graph {
	file, _ := ioutil.ReadFile(filename)
	// first line is number of edges;
	edgeLines := strings.Split(string(file), "\n")
	nodeMax, err := Atoi(edgeLines[0])
	if err != nil {
		log.Fatal("Couldn't find max number of nodes on line 1 of ", filename)
	}
	g := Graph{Edges: make([][]Edge, nodeMax), switchCost: switchCost, nodeMax: nodeMax}
	for i, _ := range g.Edges {
		// println("la")
		g.Edges[i] = make([]Edge, 0)
	}

	for _, line := range edgeLines[1 : len(edgeLines)-1] {
		//~ println("here")
		if line == "" {
			continue
		}
		var args = strings.Split(line, ",")
		start, _ := Atoi(args[0])
		end, _ := Atoi(args[1])
		cost, _ := Atoi(args[2])
		lines := strings.Split(args[3], " ")
		//fmt.Println(g)
		g.Edges[start] = append(g.Edges[int(start)], Edge{int(end), int(cost), lines})
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
	pq := make(PriorityQueue, 0) //, len(g.Edges)*100
	pq.AddNode(&Item{start, 0, 0, nil})
	distances := make([]int, len(g.Edges))

	for len(pq) > 0 {
		n := heap.Pop(&pq).(*Item)
		if visited[n.value] {
			continue
		}
		if n.value == end {
			return g.backtrack(n)
		}
		visited[n.value] = true
		distances[n.value] = n.priority
		for _, edge := range g.Edges[n.value] {
			switchCost := 0
			// switchCost, linesTraveled := g.calcSwitchCost(n, n.parent)
			// _, newlines := match(linesTraveled, edge.lines)
			// pq.AddNode(&Item{edge.end, switchCost + edge.cost + n.priority, 0, newlines, n})
			pq.AddNode(&Item{edge.end, switchCost + edge.cost + n.priority, 0, n})
		}
	}
	log.Fatal("No path between", start, "and", end, "found")
	return []int{}
}

/* function to include the switching cost of lines between nodes */
func (g *Graph) calcSwitchCost(n, p *Item) (int, []string) {
	// if n.lines != nil && n.parent != nil && n.parent.lines != n.lines {
	// 	switchCost = g.switchCost
	// }

	// if n.parent == nil || n.lines == nil || p.lines == nil {
	// 	return 0, n.lines
	// }
	// wasSwitch, travelled := match(n.lines, p.lines)
	// if wasSwitch {
	// 	return g.switchCost, travelled
	// }
	// return 0, travelled
	return 0, []string{}
}

/* function to reduce the number of lines to the least possible matches between two points; defaults to newlines */
func match(newlines []string, oldlines []string) (bool, []string) {
	wasSwitched := true
	updatedlines := make([]string, 0)
	for i := 0; i < len(oldlines); i++ {
		for j := 0; j < len(newlines); j++ {
			if oldlines[i] == newlines[j] {
				wasSwitched = false
				updatedlines = append(updatedlines, newlines[j])
			}
		}
	}
	// if there are no matches, return the new lines
	if len(updatedlines) == 0 {
		updatedlines = newlines
	}
	return wasSwitched, updatedlines
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
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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
