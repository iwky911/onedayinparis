package main

import (
	"flag"
	"fmt"
)

func buildMatrix(g Graph) {
	// fmt.Println(g.Edges)
	SaveMatrixToFile("matrix.txt", g.Matrix())
	fmt.Println("Matrix built")
}

func shortestPath(g Graph) {
	nodes := ParseNodePath("nodes.txt")
	SavePathToFile("complete_nodes.txt", g.FullPath(nodes))
	fmt.Println("Nodes written to complete_nodes.txt")
}

func main() {
	var runParseNodes = flag.Bool("p", false, "parse the nodes.txt file for the complete path")
	var switchCost = flag.Int("s", 0, "switching cost (defaults to 0)")
	flag.Parse()
	println("Switch cost set to", *switchCost)

	g := LoadGraph("edges.csv", *switchCost)

	if !*runParseNodes {
		buildMatrix(g)
	} else {
		shortestPath(g)
	}

}
