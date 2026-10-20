package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"syscall"
)

func main() {
	cmdArgs := os.Args

	if len(cmdArgs) == 1 {
		arg1 := "."
		defaultLs(arg1)
	} else {
		arg1 := cmdArgs[1]
		arg1, _ = strings.CutPrefix(arg1, "-")
		switch arg1 {
		case "a":
			listAll(checkArg(arg1), false)

		case "A":
			listAll(checkArg(arg1), true)

		case "l":
			listLongFormat(checkArg(arg1))

		default:
			defaultLs(arg1)
		}
	}

	/*
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
	*/
}

func defaultLs(dir string) {
	files, err := os.ReadDir(dir)
	if err == nil {
		for cnt := 0; cnt < len(files); cnt++ {
			filename, isHidden := strings.CutPrefix(files[cnt].Name(), ".")
			if !isHidden {
				fmt.Print(filename)
				fmt.Print("     ")
			}

		}
	} else {
		log.Fatal(err)
	}
	fmt.Println(" ")
}

func checkArg(arg1 string) string {
	if len(os.Args) == 3 {
		return os.Args[2]
	} else {
		return "."
	}
}

func listAll(dir string, almost bool) {
	files, err := os.ReadDir(dir)

	if !almost {
		fmt.Print(".     ..     ")
	}

	if err == nil {
		for cnt := 0; cnt < len(files); cnt++ {
			fmt.Print(files[cnt].Name())
			fmt.Print("     ")
		}
	} else {
		log.Fatal(err)
	}
	fmt.Println(" ")
}

func listLongFormat(dir string) {
	files, err := os.ReadDir(dir)

	if err == nil {
		for cnt := 0; cnt < len(files); cnt++ {
			filename, isHidden := strings.CutPrefix(files[cnt].Name(), ".")
			if !isHidden {
				fileinfo, _ := files[cnt].Info()
				fmt.Print(fileinfo.Mode().Perm())
				fmt.Print("     ")
				stat := fileinfo.Sys().(*syscall.Stat_t)
				userElement, _ := user.LookupId(fmt.Sprint(stat.Uid))
				groupElement, _ := user.LookupGroupId(fmt.Sprint(stat.Gid))
				fmt.Print(userElement.Username)
				fmt.Print("     ")
				fmt.Print(groupElement.Name)
				fmt.Print("     ")
				fmt.Print(fileinfo.Size())
				fmt.Print("     ")
				fmt.Print(fileinfo.ModTime().Format("02.01 15:04"))
				fmt.Print("     ")
				fmt.Print(filename)
				fmt.Println("m")
			}
		}
	} else {
		log.Fatal(err)
	}
	fmt.Println(" ")
}
