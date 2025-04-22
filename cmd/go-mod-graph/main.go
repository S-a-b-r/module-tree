package main

import (
	"flag"
	"fmt"
	"os"

	"graph-generator/internal/generator"
)

func main() {
	var depth int
	var substring string
	flag.IntVar(&depth, "depth", 5, "Глубина вложенных зависимостей")
	flag.StringVar(&substring, "substring", "", "Подстрока для поиска")
	flag.Parse()

	t := generator.GenerateTree("", substring, depth)

	fmt.Println("success get tree")
	fmt.Println(t)
	
	xml := t.ToDrawIO()

	// Сохраняем в файл
	if err := os.WriteFile("dependencies.drawio", []byte(xml), 0644); err != nil {
		fmt.Println("Error saving file:", err)
	}
}
