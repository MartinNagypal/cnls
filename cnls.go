package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"syscall"
	"text/tabwriter"
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
}

func defaultLs(dir string) {
	files, err := os.ReadDir(dir)
	sortedFiles := sort(files)
	if err == nil {
		for cnt := 0; cnt < len(sortedFiles); cnt++ {
			filename, isHidden := strings.CutPrefix(sortedFiles[cnt].Name(), ".")
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
	sortedFiles := sort(files)

	if !almost {
		fmt.Print(".     ..     ")
	}

	if err == nil {
		for cnt := 0; cnt < len(sortedFiles); cnt++ {
			fmt.Print(sortedFiles[cnt].Name())
			fmt.Print("     ")
		}
	} else {
		log.Fatal(err)
	}
	fmt.Println(" ")
}

func listLongFormat(dir string) {
	files, err := os.ReadDir(dir)
	sortedFiles := sort(files)
	table := tabwriter.NewWriter(os.Stdout, 0, 5, 5, ' ', 0)

	if err == nil {
		fmt.Println(" ")
		for cnt := 0; cnt < len(sortedFiles); cnt++ {
			filename, isHidden := strings.CutPrefix(sortedFiles[cnt].Name(), ".")
			if !isHidden {
				fileinfo, _ := sortedFiles[cnt].Info()
				permission := fileinfo.Mode().Perm()
				stat := fileinfo.Sys().(*syscall.Stat_t)
				userElement, _ := user.LookupId(fmt.Sprint(stat.Uid))
				userName := userElement.Username
				groupElement, _ := user.LookupGroupId(fmt.Sprint(stat.Gid))
				groupName := groupElement.Name
				modtime := fileinfo.ModTime().Format("02.01 15:04")
				size := fileinfo.Size()
				fmt.Fprintf(table, "%s\t%s\t%s\t%s\t%d\t%s\n", permission, userName, groupName, modtime, size, filename)
			}
		}
		table.Flush()
	} else {
		log.Fatal(err)
	}
	fmt.Println(" ")
}
