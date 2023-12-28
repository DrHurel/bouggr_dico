package dico

import (
	"bufio"
	"bytes"
	"os"
)

func RemoveFromTxt(path string, dices Dices, rp RemovePatern, lmn map[string]int) {

	var bs []byte
	var text string

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	buf := bytes.NewBuffer(bs)
	lmncount := make(map[string]int, len(lmn))
	for scanner.Scan() {
		possible := true
		for key := range lmn {
			lmncount[key] = 0
		}
		text = scanner.Text()
		for _, v := range text {
			lmncount[string(v)]++
			possible = lmn[string(v)] >= lmncount[string(v)] && possible

			for _, target := range rp[string(v)] {
				for i, t := range target {
					if i < lmncount[string(v)] {
						lmncount[t]++
						possible = lmn[t] >= lmncount[t] && possible
					} else {
						break
					}
				}
			}

		}

		length := len(text)

		if length <= 16 && length > 2 && possible {
			_, err := buf.WriteString(text + "\n")
			if err != nil {
				panic("Couldn't replace line")
			}
		}
	}
	f.Truncate(0)
	f.Seek(0, 0)
	buf.WriteTo(f)

}
