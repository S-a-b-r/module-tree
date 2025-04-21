package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"graph-generator/internal/generator"
)

func main() {
	if err := filepath.Walk(".", finderGlobarsFolders); err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}
}

func finderGlobarsFolders(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return nil
	}

	if path == "." {
		return nil
	}

	if !strings.Contains(info.Name(), "bg-") {
		return filepath.SkipDir
	}

	fmt.Printf("Found bg- directory: %s\n", path)

	generator.GenerateTree(path, "gitlab.globars.ru", 2)

	// Поиск файла go.mod в папке проекта
	// if err = filepath.Walk(path, finderGoModFilesAndRead); err != nil {
	// 	fmt.Printf("Error walking directory: %v\n", err)
	// 	os.Exit(1)
	// }

	return nil

}

// func finderGoModFilesAndRead(path string, info os.FileInfo, err error) error {
// if err != nil {
// 	return err
// }
//
// if info.IsDir() {
// 	return nil
// }
//
// if info.Name() != "go.mod" {
// 	return nil
// }
//
// fmt.Printf("Found go.mod file: %s\n", path)
//
// return printFile(path)
// }

// func printFile(path string) error {
// 	generator.GenerateTree(path, "gitlab.globars.ru", 5)
//
// 	return nil
// }
