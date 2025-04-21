package generator

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"graph-generator/internal/models"
)

func GenerateTree(dir string, substring string, depth int) {
	cmd := exec.Command("go", "mod", "graph")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Ошибка выполнения команды: %v\nВывод: %s, папка: %s", err, output, dir)
		return
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

	fmt.Println("success get tree")
	fmt.Println(g)

	xml := g.ToDrawIO()

	// Сохраняем в файл
	if err = os.WriteFile("dependencies.drawio", []byte(xml), 0644); err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("DrawIO file saved successfully")
}
