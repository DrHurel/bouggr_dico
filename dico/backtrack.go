package dico

import (
	"data_structure"
	"sync"
	"utils"
)

// Can't be used because of complexity
// func PaintGraph[T comparable, K any](G *data_structure.Node[T, K], word []T, dices []T, nbFaces int, used []bool) {

// 	if !G.CanCreateWord(word) {
// 		return
// 	} else {
// 		G.Mark(word)
// 	}

// 	for i, l := range dices {
// 		if !used[i] {
// 			for j := 0; j < nbFaces; j++ {
// 				used[i-i%nbFaces+j] = true
// 			}

// 			PaintGraph(G, append(word, l), dices, nbFaces, used)

// 			for j := 0; j < nbFaces; j++ {
// 				used[i-i%nbFaces+j] = false
// 			}
// 		}
// 	}
// }

func AllWordInGrid[T string | rune](Grid [4][4]T, dico *data_structure.Node[rune, int]) []string {

	ch := make(chan string)
	res := make([]string, 0)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		allWordInGrid(Grid, dico, ch)
	}(wg)

	for v := range ch {
		if !utils.Contain(res, v) && v != "" {
			res = append(res, v)
		}
	}

	wg.Wait()

	return res
}

func allWordInGrid[T string | rune](Grid [4][4]T, dico *data_structure.Node[rune, int], ch chan string) {

	wg := new(sync.WaitGroup)
	for i := range Grid {
		for j := range Grid {
			used := [4][4]bool{}
			wg.Add(1)
			go func(wg *sync.WaitGroup, i, j int) {
				defer wg.Done()

				appendAllWordFromGrid(
					ch,

					Grid,
					dico,
					"",
					i, j, used)
			}(wg, i, j)
		}
	}

	wg.Wait()

	close(ch)
	return
}

func appendAllWordFromGrid[T string | rune](res chan string,
	G [4][4]T,
	dico *data_structure.Node[rune, int],
	word string, i, j int, used [4][4]bool) {

	if dico.CheckWord([]rune(word), 1, func(a, b int) bool { return a == b }) {
		if len(word) != 0 {
			res <- word
		}
	}

	if !canPickAletter(i, j, used) || !dico.CanCreateWord([]rune(word)) {

		return
	}

	for _, a := range []int{-1, 0, 1} {
		for _, b := range []int{-1, 0, 1} {
			possitive := (i+a >= 0 && j+b >= 0)
			notOutOfRange := (i+a < 4 && j+b < 4)
			if possitive && notOutOfRange && !used[i+a][j+b] {
				used[i+a][j+b] = true
				word += string(G[i+a][j+b])

				appendAllWordFromGrid(res, G, dico, word, i+a, j+b, used)
				word = word[:len(word)-1]
				used[i+a][j+b] = false
			}
		}
	}

}

func canPickAletter(i int, j int, used [4][4]bool) bool {
	possible := false
	for _, a := range []int{-1, 1} {
		for _, b := range []int{-1, 1} {
			if i+a >= 0 && j+b >= 0 && i+a < 4 && j+b < 4 {
				possible = possible || !used[i+a][j+b]
			}
		}
	}
	return possible
}
