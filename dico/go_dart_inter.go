package dico

import (
	"sync"
)

func GetAllWord(grid string, dico []interface{}) []string {
	buf := make(chan string, 1024)
	resMap := make(map[string]bool)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go startAtAllPoint(buf, grid, dico, wg)

	for e := range buf {
		resMap[e] = true
	}
	res := make([]string, 0, len(resMap))
	for k := range resMap {
		res = append(res, k)
	}

	wg.Wait()

	return res
}

func startAtAllPoint(buf chan string, grid string, dico []interface{}, wg *sync.WaitGroup) {
	iwg := new(sync.WaitGroup)
	iwg.Add(16)

	for _, i := range [4]int{0, 1, 2, 3} {
		for _, j := range [4]int{0, 1, 2, 3} {

			go initPoint(buf, grid, iwg, i, j, dico,
				grid[i*4+j])
		}
	}

	iwg.Wait()
	close(buf)
	defer wg.Done()

}

func initPoint(buf chan string, grid string, iwg *sync.WaitGroup, i, j int, dico []interface{}, point byte) {
	defer iwg.Done()
	var used [4][4]bool
	used[i][j] = true
	for _, n := range dico[1].([]interface{}) {
		child := n.([]interface{})
		if (int(child[0].(int32)) & int(point)) == int(point) {
			appendFromPoint(buf, grid, string(point), i, j, used, child)
		}
	}
}

func appendFromPoint(res chan string, grid string, word string, i, j int, used [4][4]bool, node []interface{}) {
	n := len(node)
	var value int = int(node[0].(int32))

	if (value & (1 << 8)) > 0 {
		res <- word
	}
	if n != 2 {
		return
	}
	var ix, jy, index int

	children := node[1].([]interface{})

	for _, a := range []int{-1, 0, 1} {
		for _, b := range []int{-1, 0, 1} {
			ix = a + i
			jy = b + j
			index = ix*4 + jy
			if ix < 0 || jy < 0 || ix > 3 || jy > 3 {
				continue
			}
			if used[ix][jy] {
				continue
			}

			newLetter := grid[index]
			for _, n := range children {
				child, ok := n.([]interface{})

				if !ok && byte(n.(int32)) == byte(newLetter) {
					used[ix][jy] = true
					appendFromPoint(res, grid, word+string(newLetter), ix, jy, used, []interface{}{n})
					used[ix][jy] = false
					break
				}

				if !ok {
					continue
				}

				if byte(child[0].(int32)) == byte(newLetter) {
					used[ix][jy] = true
					appendFromPoint(res, grid, word+string(newLetter), ix, jy, used, child)
					used[ix][jy] = false
					break
				}

			}
		}
	}
}

func sendToChan(node [2]interface{}, res chan string, word string) {

	if val, ok := node[0].(int32); ok {

		if (int32(val) & (1 << 8)) > 0 {
			//log.Println(node[0], word, string(byte(val)))
			res <- word
		}
	}
}
