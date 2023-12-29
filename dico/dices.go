package dico

type Dices [16]Dice
type LetterOptionMap map[string][]string
type RemovePatern map[string][][]string
type IterationCount map[string]int

func checkFace(l rune, lomap LetterOptionMap, dice Dice) {
	check := make(map[rune]bool)
	for _, face := range dice {
		if !check[face] && face != l {
			lomap[string(l)] = append(lomap[string(l)], string(face))
			check[face] = true
		}
	}

}

// The `LetterOption` function is a method of the `Dices` type. It calculates the possible options for
// each letter in the dice set.
func (this Dices) LetterOption() LetterOptionMap {
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

func (this Dices) RemovePaternStruct(lmn IterationCount, lom LetterOptionMap) RemovePatern {
	allLetter := "ABCDEFGHIJKLMNOPQRSTUVWYXZ"

	res := make(RemovePatern)
	for _, letter := range allLetter {
		n := lmn[string(letter)]
		res[string(letter)] = make([][]string, n)
		for _, letter2 := range allLetter {
			if letter2 != letter {
				if count := numberOfIteration(lom[string(letter)], letter2); count != 0 {
					res[string(letter)][n-count] = append(res[string(letter)][n-count], string(letter2))
				}
			}

		}

	}

	return res

}

func (this Dices) LetterMaxNumber() IterationCount {

	res := make(IterationCount)
	for _, dice := range this {
		check := make(map[rune]bool)
		for _, face := range dice {
			if !check[face] {
				check[face] = true
				res[string(face)]++
			}
		}
	}

	return res
}
