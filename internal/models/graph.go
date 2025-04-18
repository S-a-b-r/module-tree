package models

import (
	"fmt"
)

type Graph struct {
	*Module
}

func NewGraph(graphStrs []string) *Graph {
	nameGraph := graphStrs[0]

	g := Graph{
		&Module{
			Name:    nameGraph,
			Version: "1.0.0",
		},
	}

	for i := 0; i+1 < len(graphStrs); i += 2 {
		parentName, _ := parseModuleInfo(graphStrs[i])
		childName, childVersion := parseModuleInfo(graphStrs[i+1])

		if !g.AddModule(parentName, childName, childVersion) {
			return nil
		}
	}

	return &g
}

func (g *Graph) String() string {
	if g.Module == nil {
		return "<empty graph>"
	}
	return g.Module.String(0, "")
}

func (g *Graph) GetWithDepth(depth int) *Graph {
	if depth < 2 {
		fmt.Println("Min depth is 2")
		return nil
	}

	newG := Graph{
		&Module{
			Name:         g.Name,
			Version:      g.Version,
			ChildModules: make([]*Module, 0),
		},
	}

	for _, m := range g.ChildModules {
		newG.ChildModules = append(newG.ChildModules, m.GetWithDepth(depth-1))
	}
	return &newG
}
