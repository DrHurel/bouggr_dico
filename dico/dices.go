package dico

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Dices [16]Dice
type LetterOptionMap map[string][]string
type RemovePatern map[string][][]string

func checkFace(l rune, lomap LetterOptionMap, dice Dice) {
	check := make(map[rune]bool)
	for _, face := range dice {
		if !check[face] && face != l {
			lomap[string(l)] = append(lomap[string(l)], string(face))
			check[face] = true
		}
	}
}

func (this Dices) LetterOption(path ...string) LetterOptionMap {
	allLetter := "ABCDEFGHIJKLMNOPQRSTUVWYXZ"
	print(len(allLetter))

	res := make(LetterOptionMap, 0)

	for _, l := range allLetter {

		for _, dice := range this {
			if dice.Contains(l) {
				checkFace(l, res, dice)
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

func numberOfIteration(LetterOption []string, target rune) int {
	count := 0
	for _, l := range LetterOption {
		if l == string(target) {
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
		n := letterMaxNumber[string(letter)]
		res[string(letter)] = make([][]string, n)
		for _, letter2 := range allLetter {
			if letter2 != letter {
				if count := numberOfIteration(LetterOptionMap[string(letter)], letter2); count != 0 {
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

func (this Dices) LetterMaxNumber(path ...string) map[string]int {

	res := make(map[string]int)
	for _, dice := range this {
		check := make(map[rune]bool)
		for _, face := range dice {
			if !check[face] {
				check[face] = true
				res[string(face)]++
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
