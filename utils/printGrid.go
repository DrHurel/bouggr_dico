package utils

import "fmt"

func PrintGrid(Grid [4][4]rune) {
	for _, e := range Grid {
		for _, c := range e {
			fmt.Print(string(c))
		}
		fmt.Print("\n")
	}

}
