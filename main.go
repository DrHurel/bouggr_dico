package main

import (
	"dico"
	"log"
	"os"
	"time"
	"tree"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("no input and ouput")
	}

	input := os.Args[1]
	output := os.Args[2]

	log.Print("Generate\n")
	start := time.Now()

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

	rp := dices.RemovePaternStruct(lmn, lom)

	dico.RemoveOfTxt(input, dices, rp, lmn)

	origin := tree.GenerateDicoFromTxt(input)
	origin.Encode(output)

	elapsed := time.Since(start)
	log.Printf("Took %s\n", elapsed)
}
