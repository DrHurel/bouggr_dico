package dico

import (
	"bufio"
	"bytes"
	"os"
)

func isPossible(word string, interationCount IterationCount, maxAllowed IterationCount, rp RemovePatern) bool {
	possible := true
	for _, v := range word {

		interationCount[string(v)]++ //ce que j'ai vraiment
		// on a forcément la lettre dans le mots
		possible = maxAllowed[string(v)] >= interationCount[string(v)] && possible

		if !possible { //une fois que possible est faux elle le reste : reduit la complexité en moyenne
			return false
		}

		for _, target := range rp[string(v)] {
			for i, t := range target {

				// garanti qu'on ne supprise pas des options non voulu
				if i >= interationCount[string(v)] {
					break
				}

				maxAllowed[t]-- //ce que je suis autorisé à avoir

				//le nombre de lettre max est sup au nombre de lettre dans le mot ou on a pas la lettre dans le mot
				possible = (maxAllowed[t] >= interationCount[t] || interationCount[t] == 0) && possible

			}
		}

	}

	//on renvoit pas vrai pour le cas où la dernière lettre rend possible faux
	return possible

}

// The function `RemoveFromTxt` reads a text file, applies certain patterns and conditions to each
// line, and writes the lines that meet the conditions back to the file.

func RemoveOfTxt(path string, dices Dices, rp RemovePatern, lmn IterationCount) {

	var bs []byte
	var text string
	var length int
	interationCount := make(IterationCount, len(lmn))
	maxAllowed := make(IterationCount, len(lmn))

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
			interationCount[key] = 0 //nombre d'iteration d'une lettre dans un mot
			maxAllowed[key] = val    //copy des bornes sups
		}

		//determination de la validité d'un mot
		possible := isPossible(text, interationCount, maxAllowed, rp)

		// garde les mots valides
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
