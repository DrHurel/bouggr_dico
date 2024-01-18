package dico

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"ui"
)

type Dices []Dice
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

func GetDicesFromTxt(path string) Dices {
	res := Dices{}

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Print(err, '\n')
		ui.PrintHelp()
		os.Exit(0)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		temp := Dice{}
		for _, l := range text {
			temp = append(temp, l)
		}
		res = append(res, temp)

	}

	return res
}

func (dices Dices) Roll(r *rand.Rand) [4][4]rune {
	var grid [4][4]rune

	r.Shuffle(len(dices), func(i, j int) {
		dices[i], dices[j] = dices[j], dices[i]
	})
	k := 0
	for i := range grid {
		for j := range grid {

			faces := r.Intn(5)

			grid[i][j] = dices[k][faces]
			k++
		}
	}
	return grid
}
