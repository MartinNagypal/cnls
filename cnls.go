package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/user"
	"strings"
	"syscall"
	"text/tabwriter"

	"golang.org/x/term"
)

const COL_WIDTH int = 2

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

func checkArg(arg1 string) string {
	if len(os.Args) == 3 {
		return os.Args[2]
	} else {
		return "."
	}
}

func getTerminalWidth() int {
	width, _, termSizeErr := term.GetSize(int(os.Stdout.Fd()))
	if termSizeErr != nil {
		log.Fatal(termSizeErr)
	}

	return width
}

func getMaxLen(sliceToCount []os.DirEntry, countHidden bool) int {
	maxlen := 0
	for _, file := range sliceToCount {
		if countHidden {
			if len(file.Name()) > maxlen {
				maxlen = len(file.Name())
			}
		} else {
			if !strings.HasPrefix(file.Name(), ".") && len(file.Name()) > maxlen {
				maxlen = len(file.Name())
			}
		}

	}
	return maxlen
}

func defaultLs(dir string) {
	fmt.Println()
	files, err := os.ReadDir(dir) //dateien nehmen
	sortedFiles := sort(files)    //sortieren
	//err handling
	if err != nil {
		log.Fatal(err)
	}

	//formatierung
	maxlen := getMaxLen(sortedFiles, false)
	cols := int(math.Floor(float64(getTerminalWidth() / (maxlen + (COL_WIDTH * 2)))))
	colWidth := maxlen + COL_WIDTH
	cnt := 0

	for _, file := range sortedFiles {
		if !strings.HasPrefix(file.Name(), ".") {
			fmt.Printf("%-*s", colWidth, file.Name())
			cnt++
		}
		if cnt == cols {
			fmt.Println()
			cnt = 0
		}
	}
	fmt.Println()
	fmt.Println()
}

func listAll(dir string, almost bool) {
	fmt.Println()
	files, err := os.ReadDir(dir)
	sortedFiles := sort(files)

	if err != nil {
		log.Fatal(err)
	}

	maxlen := getMaxLen(sortedFiles, false)
	cols := int(math.Floor(float64(getTerminalWidth() / (maxlen + (COL_WIDTH * 2)))))
	colWidth := maxlen + COL_WIDTH
	cnt := 0

	if !almost {
		dots := []string{".", ".."}
		for _, element := range dots {
			fmt.Printf("%-*s", colWidth, element)
			cnt++
			if cnt == cols {
				fmt.Println()
				cnt = 0
			}
		}
	}

	for _, file := range sortedFiles {
		fmt.Printf("%-*s", colWidth, file.Name())
		cnt++

		if cnt == cols {
			fmt.Println()
			cnt = 0
		}
	}
	fmt.Println()
	fmt.Println()
}

func listLongFormat(dir string) {
	fmt.Println()
	files, err := os.ReadDir(dir)
	sortedFiles := sort(files)
	table := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)

	if err == nil {
		fmt.Println(" ")
		fmt.Fprintf(table, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "perm", "d/f", "owner", "group", "last modified", "size", "filename")
		for cnt := 0; cnt < len(sortedFiles); cnt++ {
			filename, isHidden := strings.CutPrefix(sortedFiles[cnt].Name(), ".")
			if !isHidden {
				fileinfo, _ := sortedFiles[cnt].Info()
				permission := fileinfo.Mode().Perm()
				isDirectory := fileinfo.Mode().IsDir()
				var isDirectoryString string
				if isDirectory {
					isDirectoryString = "d"
				} else {
					isDirectoryString = "f"
				}
				stat := fileinfo.Sys().(*syscall.Stat_t)
				userElement, _ := user.LookupId(fmt.Sprint(stat.Uid))
				userName := userElement.Username
				groupElement, _ := user.LookupGroupId(fmt.Sprint(stat.Gid))
				groupName := groupElement.Name
				modtime := fileinfo.ModTime().Format("02.01 15:04")
				size := fileinfo.Size()
				fmt.Fprintf(table, "%s\t%s\t%s\t%s\t%s\t%d\t%s\n", permission, isDirectoryString, userName, groupName, modtime, size, filename)
			}
		}
		table.Flush()
	} else {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println()
}
