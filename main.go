package main

import (
	"bufio"
	"dico"
	"log"
	"os"
	"tree"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("no input and ouput")
	}

	input := os.Args[1]
	output := os.Args[2]

	log.Print("Generate\n")
	origin := new(tree.Node[string])
	origin.Key = "origin"

	dices := dico.Dices{
		dico.Dice{'E', 'T', 'U', 'K', 'N', 'O'},
		dico.Dice{'E', 'V', 'G', 'T', 'I', 'N'},
		dico.Dice{'D', 'E', 'C', 'A', 'M', 'P'},
		dico.Dice{'I', 'E', 'L', 'R', 'U', 'W'},
		dico.Dice{'E', 'H', 'I', 'F', 'S', 'E'},
		dico.Dice{'R', 'E', 'C', 'A', 'L', 'S'},
		dico.Dice{'E', 'N', 'T', 'D', 'O', 'S'},
		dico.Dice{'O', 'F', 'X', 'R', 'I', 'A'},
		dico.Dice{'N', 'A', 'V', 'E', 'D', 'Z'},
		dico.Dice{'E', 'I', 'O', 'A', 'T', 'A'},
		dico.Dice{'G', 'L', 'E', 'N', 'Y', 'U'},
		dico.Dice{'B', 'M', 'A', 'Q', 'J', 'O'},
		dico.Dice{'T', 'L', 'I', 'B', 'R', 'A'},
		dico.Dice{'S', 'P', 'U', 'L', 'T', 'E'},
		dico.Dice{'A', 'I', 'M', 'S', 'O', 'R'},
		dico.Dice{'E', 'N', 'H', 'R', 'I', 'S'},
	}

	lmn := dices.LetterMaxNumber()
	lom := dices.LetterOption()

	rp := dices.RemoveIfPick(lmn, lom)

	//TODO : Etape 3 Suppression de mots pas réalisable

	dico.RemoveFromTxt(input, dices, rp, lmn)

	//TODO : Etape 4 construction de l'arbre dictionnaire

	f, err := os.OpenFile(input, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	//var bs []byte
	//buf := bytes.NewBuffer(bs)

	var text string

	for scanner.Scan() {
		text = scanner.Text()

		next := origin

		for _, l := range text {
			if temp, err := next.HasChild(string(l)); err == nil {
				next = temp
			} else {
				temp = new(tree.Node[string])
				temp.Childs = []*tree.Node[string]{}
				temp.Key = string(l)
				next.Add(temp)
				next = temp
			}
		}
		next.Valide = 1

	}

	origin.Encode(output)

	log.Print("Done\n")
}
