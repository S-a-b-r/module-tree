package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"graph-generator/internal/models"
)

func main() {
	var depth int
	flag.IntVar(&depth, "depth", 5, "Глубина вложенных зависимостей")
	flag.Parse()

	cmd := exec.Command("go", "mod", "graph")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Ошибка выполнения команды: %v\nВывод: %s", err, output)
	}

	strsModules := strings.Split(string(output), "\n")
	strs := make([]string, 0)
	for _, s := range strsModules {
		ss := strings.Split(s, " ")
		strs = append(strs, ss...)
	}
	g := models.NewGraph(strs)

	fmt.Println(g.GetWithDepth(depth))
}
