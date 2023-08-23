package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

const gvFileName = "go-mod-graph.gv"

func main() {
	cmd := exec.Command("go", "mod", "graph")
	out, err := cmd.Output()
	CheckError(err)
	lines := strings.Split(string(out), "\n")
	CheckError(err)

	g := graph.New(graph.StringHash, graph.Directed())
	
	for i := range lines {
		nodes := strings.Fields(lines[i])
		//TODO remove version number if -v flag not given
		for i := range nodes {
			nodes[i] = strings.Split(nodes[i], "@")[0]
		}
		if len(nodes) != 2 {continue}

		_ = g.AddVertex(nodes[0])
		_ = g.AddVertex(nodes[1])
		_ = g.AddEdge(nodes[0], nodes[1])
	}

	file, _ := os.Create(gvFileName)
	err = draw.DOT(g, file,
		draw.GraphAttribute("overlap", "false"),
		draw.GraphAttribute("pack", "true"),
		draw.GraphAttribute("normalize", "true"),
	)
	CheckError(err)
	//TODO only open the svg if -o flag is given
	//neato puts the current project's node in the centre
	cmd = exec.Command("dot", "-Tsvg", "-Kneato", "-O", gvFileName)
	out, err = cmd.Output()
	CheckError(err)
	cmd = exec.Command("open", gvFileName + ".svg")
	out, err = cmd.Output()
	CheckError(err)
	
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Could not produce graph: ", err)
		os.Exit(1)

	}
}
