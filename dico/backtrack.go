package dico

import (
	"data_structure"
	"sync"
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

func AllWordInGrid[T rune](Grid [4][4]T, dico *data_structure.Node, lang int32) []string {

	ch := make(chan string, 1024)
	resMap := make(map[string]bool)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		allWordInGrid(Grid, dico, ch, lang)
	}(wg)

	for v := range ch {
		resMap[v] = true
	}

	res := make([]string, 0, len(resMap))
	for k := range resMap {
		res = append(res, k)
	}

	wg.Wait()

	return res
}

func allWordInGrid[T rune](Grid [4][4]T, dico *data_structure.Node, ch chan string, lang int32) {

	wg := new(sync.WaitGroup)
	for i := range Grid {
		for j := range Grid {
			used := [4][4]bool{}
			used[i][j] = true
			wg.Add(1)
			go func(wg *sync.WaitGroup, i, j int) {
				defer wg.Done()
				d, _ := dico.GetChild(rune(Grid[i][j]))
				appendAllWordFromPoint(
					ch,

					Grid,
					d,
					string(Grid[i][j]),
					i, j, used, lang)
			}(wg, i, j)
		}
	}

	wg.Wait()

	close(ch)
	return
}

func appendAllWordFromPoint[T rune](res chan string,
	G [4][4]T,
	dico *data_structure.Node,
	word string, i, j int,
	used [4][4]bool,
	lang int32) {

	if dico.Value&data_structure.IS_A_WORD > 0 {
		if len(word) != 0 {
			res <- word
		}
	}
	var ix, jy int

	for _, a := range []int{-1, 0, 1} {
		for _, b := range []int{-1, 0, 1} {
			ix = a + i
			jy = b + j
			if !(ix > -1 && jy > -1) {
				continue
			}
			if !(ix < 4 && jy < 4) {
				continue
			}
			if !used[ix][jy] {
				used[ix][jy] = true
				if v, ok := dico.GetChild(rune(G[ix][jy])); ok == nil {
					appendAllWordFromPoint(res, G, v, word+string(G[ix][jy]), ix, jy, used, lang)
				}

				used[ix][jy] = false
			}
		}
	}

}

func canPickAletter(i int, j int, used [4][4]bool) bool {
	possible := false
	for _, a := range []int{-1, 0, 1} {
		for _, b := range []int{-1, 0, 1} {
			if i+a >= 0 && j+b >= 0 && i+a < 4 && j+b < 4 {
				possible = possible || !used[i+a][j+b]
			}
		}
	}
	return possible
}
