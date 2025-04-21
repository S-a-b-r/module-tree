package lib

import (
	"fmt"
	"io/ioutil"
)

func ListDirByReadDir(path string) {

	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("[%s]\n", val.Name())
		} else {
			fmt.Println(val.Name())
		}
	}
}
