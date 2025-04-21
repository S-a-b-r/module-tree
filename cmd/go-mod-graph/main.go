package main

import (
	"flag"

	"graph-generator/internal/generator"
)

func main() {
	var depth int
	var substring string
	flag.IntVar(&depth, "depth", 5, "Глубина вложенных зависимостей")
	flag.StringVar(&substring, "substring", "", "Подстрока для поиска")
	flag.Parse()

	generator.GenerateTree("", substring, depth)
}
