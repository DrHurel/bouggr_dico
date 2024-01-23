package data_structure

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"ui"
)

const ORIGIN_KEY = 'O'

// The function `GenerateFromTxt` reads a text file, creates a tree structure based on the characters
// in the file, and returns the root node of the tree.
func GenerateDicoFromTxt(input string, lang map[string]int16) *Node {

	origin := new(Node)

	var currentLang int16
	origin.Value = ORIGIN_KEY

	f, err := os.OpenFile(input, os.O_RDWR, 0644)
	if err != nil {
		fmt.Print(err, '\n')
		ui.PrintHelp()
		os.Exit(0)
	}

	scanner := bufio.NewScanner(f)

	var text string

	for scanner.Scan() {
		text = scanner.Text()

		next := origin
		if len(text) > 2 && text[1] == '=' {
			currentLang = lang[strings.Split(text, "=")[1]]

		} else {
			for _, l := range text {
				if temp, err := next.GetChild((l)); err == nil {
					next = temp
				} else {
					temp = NewNode(int16(l))
					next.Add(temp)
					next = temp

				}
			}
			next.AddLang(currentLang)
			next.SetAWord()

		}

	}

	return origin
}
