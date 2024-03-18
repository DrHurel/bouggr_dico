package dico

import (
	"log"
	"sync"
)

func GetAllWord(grid string, dico [2]interface{}) []string {
	buf := make(chan string)
	res := make([]string, 0)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		iwg := new(sync.WaitGroup)
		iwg.Add(16)

		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {

				go func(iwg *sync.WaitGroup, i, j int, dico [2]interface{}) {
					defer iwg.Done()
					used := [4][4]bool{}
					used[i][j] = true
					for _, n := range dico[1].([]interface{}) {
						children, ok := n.([]interface{})
						if !ok {
							continue
						}
						if (int(children[0].(float64)) & int(grid[i*4+j])) == int(grid[i*4+j]) {
							appendFromPoint(buf, grid, string(grid[i*4+j]), i, j, used, [2]interface{}(children))
						}
					}
				}(iwg, i, j, dico)
			}
		}

		iwg.Wait()
		close(buf)
		defer wg.Done()

	}(wg)

	for e := range buf {

		if !contains(res, e) {
			res = append(res, e)
		}

	}

	wg.Wait()

	return res
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func appendFromPoint(res chan string, grid string, word string, x, y int, used [4][4]bool, node [2]interface{}) {
	if len(node) == 0 {
		log.Println("empty")
		return
	}

	if val, ok := node[0].(float64); ok { // If the node is a number (should be useless but who knows)
		if (int16(val) & (1 << 8)) > 0 {
			log.Println("word", word, string(grid[x*4+y]), val, val-(1<<8)-(1<<9))
			res <- word
		}
	}

	if len(node) == 1 { // If the node is a leaf
		return
	}
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			ix := x + i
			jy := y + j
			index := ix*4 + jy
			if ix < 0 || ix > 3 || jy < 0 || jy > 3 { // If the point is out of the grid
				continue
			}
			children, ok := node[1].([]interface{})
			if used[ix][jy] || !ok { // If the point is already used or the node is not a list
				continue
			}
			for _, n := range children {
				child, ok := n.([]interface{})
				if !ok || len(child) < 1 {
					continue
				}
				if v, ok := child[0].(float64); ok && byte(v) == byte(grid[index]) {
					used[ix][jy] = true
					appendFromPoint(res, grid, word+string(grid[index]), ix, jy, used, [2]interface{}(child))
					used[ix][jy] = false
					break
				}

			}

		}
	}
}
