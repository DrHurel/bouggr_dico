package tree

import (
	"bufio"
	"os"
)

const ORIGIN_KEY = "origin"

// The function `GenerateFromTxt` reads a text file, creates a tree structure based on the characters
// in the file, and returns the root node of the tree.
func GenerateDicoFromTxt(input string) *Node[string, int] {

	origin := new(Node[string, int])
	origin.Key = ORIGIN_KEY

	f, err := os.OpenFile(input, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	var text string

	for scanner.Scan() {
		text = scanner.Text()

		next := origin

		for _, l := range text {
			if temp, err := next.GetChild(string(l)); err == nil {
				next = temp
			} else {
				temp = NewNode[string, int](string(l))
				next.Add(temp)
				next = temp
			}
		}
		next.Value = 1

	}

	return origin
}
