package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	cmdArgs := os.Args
	var dir string
	if len(cmdArgs) != 2 {
		dir = "."
	} else {
		dir = cmdArgs[1]
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for file := 0; file < len(files); file++ {
		var filename string = files[file].Name()
		filename, startWithDot := strings.CutPrefix(filename, ".")
		fileInfo, fileInfoErr := files[file].Info()
		if !startWithDot {
			if fileInfoErr == nil {
				fmt.Print(fileInfo.Mode().Perm())
				fmt.Print("   ")
				fmt.Println(filename)
			}
		}
	}
	fmt.Println(" ")
}
