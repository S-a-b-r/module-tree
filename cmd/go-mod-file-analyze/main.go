package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"graph-generator/internal/generator"
)

func main() {
	var prefix, substring string
	var depth int
	flag.IntVar(&depth, "depth", 5, "Глубина вложенных зависимостей")
	flag.StringVar(&substring, "substring", "", "Подстрока для поиска")
	flag.StringVar(&prefix, "prefix", "", "Префикс для поиска названия папок с микросервисами")
	flag.Parse()

	fmt.Println("Start script")

	if err := filepath.Walk(".", handleFoldersWithPrefix(prefix, substring, depth)); err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("End script")
}

func handleFoldersWithPrefix(prefix, substring string, depth int) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() || !strings.Contains(info.Name(), prefix) {
			return nil
		}

		fmt.Printf("Found directory with prefix: %s\n", path)

		tree := generator.GenerateTree(path, substring, depth)
		if tree == nil {
			fmt.Printf("Failed to generate tree from %s", path)
			return nil
		}

		fmt.Println(tree)

		xml := tree.ToDrawIO()

		drawIOfile := fmt.Sprintf("drawio/%s.dependencies.drawio", info.Name())
		if err = os.WriteFile(drawIOfile, []byte(xml), 0644); err != nil {
			fmt.Println("Error saving file:", err)
		}

		return nil
	}
}
