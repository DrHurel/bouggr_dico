package main

import (
	"dico"
	"log"
)

func main() {

	log.Print("Generate\n")

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

	dices.RemoveIfPick("out", "test.json")

	//TODO : Etape 3 Suppression de mots pas réalisable

	dico.RemoveFromTxt("dico copy.txt", dices, "./out/test.json")

	//TODO : Etape 4 construction de l'arbre dictionnaire

	log.Print("Done\n")
}
