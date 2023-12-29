package dico

import (
	"bufio"
	"bytes"
	"os"
)

func isPossible(word string, lmn map[string]int, lmncount map[string]int, lmncopy map[string]int, rp RemovePatern) bool {
	possible := true
	for _, v := range word {

		lmncount[string(v)]++
		// on a forcément la lettre dans le mots
		possible = lmncopy[string(v)] >= lmncount[string(v)] && possible

		if !possible {
			return false
		}

		for _, target := range rp[string(v)] {
			for i, t := range target {
				if i < lmncount[string(v)] {
					lmncopy[t]--
					//le nombre de lettre max est sup au nombre de lettre dans le mot ou on a pas la lettre dans le mot
					possible = (lmncopy[t] >= lmncount[t] || lmncount[t] == 0) && possible
				} else {
					break
				}
			}
		}

	}

	return possible

}

// The function `RemoveFromTxt` reads a text file, applies certain patterns and conditions to each
// line, and writes the lines that meet the conditions back to the file.

func RemoveFromTxt(path string, dices Dices, rp RemovePatern, lmn map[string]int) {

	var bs []byte
	var text string
	var length int
	lmncount := make(map[string]int, len(lmn))
	lmncopy := make(map[string]int, len(lmn))

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	buf := bytes.NewBuffer(bs)

	//pour tout les lines du fichier
	for scanner.Scan() {

		//variable d'entrer
		text = scanner.Text()
		length = len(text)

		for key, val := range lmn {
			lmncount[key] = 0  //nombre d'iteration d'une lettre dans un mot
			lmncopy[key] = val //copy des bornes sups
		}

		//determination de la validité d'un mot
		possible := isPossible(text, lmn, lmncount, lmncopy, rp)
		// garde les mots valide
		if possible && length <= 16 && length > 2 {
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
