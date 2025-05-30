package models

import (
	"fmt"
	"strings"
)

type Tree struct {
	*Module
}

func NewTree(graphStrs []string) *Tree {
	nameTree := graphStrs[0]

	g := Tree{
		&Module{
			Name:    nameTree,
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

func (g *Tree) String() string {
	if g.Module == nil {
		return "<empty graph>"
	}
	return g.Module.String(0, "")
}

func (g *Tree) GetWithDepth(depth int) *Tree {
	if depth < 2 {
		fmt.Println("Min depth is 2")
		return nil
	}

	newG := Tree{
		&Module{
			Name:         g.Name,
			Version:      g.Version,
			ChildModules: make([]*Module, 0),
		},
	}

	for _, m := range g.ChildModules {
		cm := m.GetWithDepth(depth - 1)
		if cm != nil {
			newG.ChildModules = append(newG.ChildModules, cm)
		}
	}
	return &newG
}

func (g *Tree) GetWithSubstr(substr string) *Tree {
	if substr == "" {
		fmt.Println("Empty substr")
		return nil
	}

	newG := Tree{
		&Module{
			Name:         g.Name,
			Version:      g.Version,
			ChildModules: make([]*Module, 0),
		},
	}

	for _, m := range g.ChildModules {
		cm := m.GetWithSubstr(substr)
		if cm != nil {
			newG.ChildModules = append(newG.ChildModules, cm)
		}
	}

	return &newG
}

// ToDrawIO генерирует XML для импорта в draw.io
func (g *Tree) ToDrawIO() string {
	if g.Module == nil {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(`<mxfile>
	<diagram name="Page-1">
		<mxTreeModel>
			<root>
				<mxCell id="0" />
				<mxCell id="1" parent="0" />`)

	// Рекурсивное построение элементов
	nextId := 2
	g.Module.DrawIONode(&sb, 1, &nextId, 0, 0)

	sb.WriteString(`
			</root>
		</mxTreeModel>
	</diagram>
</mxfile>`)

	return sb.String()
}
