package dico

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Dices [16]Dice
type LetterOptionMap map[rune][]rune
type RemovePatern map[string][][]string

func (this Dices) LetterOption(path ...string) LetterOptionMap {
	allLetter := "ABCDEFGHIJKLMNOPQRSTUVWYXZ"
	print(len(allLetter))

	res := make(LetterOptionMap, 0)

	for _, l := range allLetter {

		for _, dice := range this {
			check := make(map[rune]bool)
			if dice.Contains(l) {
				for _, face := range dice {
					if !check[face] && face != l {
						res[l] = append(res[l], face)
						check[face] = true
					}
				}
			}
		}
	}

	if len(path) > 0 {
		file, err := os.Create(path[0])
		if err != nil {
			panic(err)
		}

		encoder := json.NewEncoder(file)
		encoder.Encode(res)
	}

	return res
}

func numberOfIteration(LetterOption []rune, target rune) int {
	count := 0
	for _, l := range LetterOption {
		if l == target {
			count++
		}
	}
	return count
}

func (this Dices) RemoveIfPick(out ...string) RemovePatern {
	allLetter := "ABCDEFGHIJKLMNOPQRSTUVWYXZ"

	path := filepath.Join(out...)
	tmp := make([]string, len(out))
	n := copy(tmp, out)
	tmp[n-1] = "lm-" + out[n-1]
	letterMaxNumber := this.LetterMaxNumber(filepath.Join(tmp...))
	tmp[n-1] = "lo-" + out[n-1]
	LetterOptionMap := this.LetterOption(filepath.Join(tmp...))
	res := make(RemovePatern, 0)
	for _, letter := range allLetter {
		n := letterMaxNumber[letter]
		res[string(letter)] = make([][]string, n)
		for _, letter2 := range allLetter {
			if letter2 != letter {
				if count := numberOfIteration(LetterOptionMap[letter], letter2); count != 0 {
					res[string(letter)][n-count] = append(res[string(letter)][n-count], string(letter2))
				}
			}

		}

	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(res)

	return res

}

func (this Dices) LetterMaxNumber(path ...string) map[rune]int {

	res := make(map[rune]int)
	for _, dice := range this {
		check := make(map[rune]bool)
		for _, face := range dice {
			if !check[face] {
				check[face] = true
				res[face]++
			}
		}
	}

	if len(path) > 0 {
		file, err := os.Create(path[0])
		if err != nil {
			panic(err)
		}

		encoder := json.NewEncoder(file)

		encoder.Encode(res)
	}

	return res
}
