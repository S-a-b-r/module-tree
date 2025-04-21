package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"graph-generator/internal/models"
)

func main() {
	var depth int
	var substring string
	flag.IntVar(&depth, "depth", 5, "Глубина вложенных зависимостей")
	flag.StringVar(&substring, "substring", "", "Подстрока для поиска")
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
