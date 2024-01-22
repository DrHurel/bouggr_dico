package data_structure

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"ui"
	"utils"
)

const ORIGIN_KEY = 'O'

// The function `GenerateFromTxt` reads a text file, creates a tree structure based on the characters
// in the file, and returns the root node of the tree.
func GenerateDicoFromTxt(input string, lang map[string]int) *Node[rune, int] {

	origin := new(Node[rune, int])

	var currentLang int
	origin.Key = ORIGIN_KEY

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
					temp = NewNode[rune, int]((l))
					next.Add(temp)
					next = temp

				}
			}
			next.Value = 1
			next.Langue += LangueCode(currentLang)
		}

	}

	return origin
}

type PendingRemovableNode[T comparable, K any] struct {
	Parent *Node[T, K]
	Index  int
}

func RemoveUnaccessible[T comparable, K any](G *Node[T, K], test K, equal func(K, K) bool) {
	visited := []*Node[T, K]{}
	p := new(Pile[*Node[T, K]])
	p.Empiler(G)
	toBeRemove := make([]PendingRemovableNode[T, K], 0)

	for p.Len() > 0 {
		G = p.Depiler()
		if !utils.Contain(visited, G) {
			visited = append(visited, G)
			G.removeUnaccessibleChildren(toBeRemove)

			for _, child := range G.Children {
				if utils.Contain(visited, child) {
					p.Empiler(child)
				}
			}
		}
	}

	for _, e := range toBeRemove {
		e.Parent.Children[e.Index] = nil
	}

}
func (this *Node[T, K]) removeUnaccessibleChildren(toBeRemove []PendingRemovableNode[T, K]) {
	for i := 0; i < len(this.Children); i++ {
		if !this.Children[i].accessible {
			toBeRemove = append(toBeRemove, PendingRemovableNode[T, K]{
				Parent: this,
				Index:  i,
			})

		}
	}
}
