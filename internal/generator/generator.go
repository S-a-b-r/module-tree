package generator

import (
	"fmt"
	"os/exec"
	"strings"

	"graph-generator/internal/models"
)

func GenerateTree(dir string, substring string, depth int) *models.Graph {
	cmd := exec.Command("go", "mod", "graph")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Ошибка выполнения команды: %v\nВывод: %s, папка: %s", err, output, dir)
		return nil
	}

	strsModules := strings.Split(string(output), "\n")
	strs := make([]string, 0)
	for _, s := range strsModules {
		ss := strings.Split(s, " ")
		strs = append(strs, ss...)
	}
	
	g := models.NewGraph(strs).GetWithDepth(depth)

	if substring != "" {
		g = g.GetWithSubstr(substring)
	}

	return g
}
