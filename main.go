package main

import (
	"dico"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"tree"
	"ui"
	"utils"
)

func main() {

	params := ui.GetParams()

	input := os.Args[params[ui.TARGET]]
	var output string
	if params[ui.SPECIAL_OUPUT] == -1 {
		output = "./out/dico.json"
	} else {
		output = strings.Split(os.Args[params[ui.SPECIAL_OUPUT]], "=")[1]
	}

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

	if params[ui.DICE_FILE] != -1 {
		dices = dico.GetDicesFromTxt(strings.Split(os.Args[params[ui.DICE_FILE]], "=")[1])
	}

	//EN
	/*
		dices = dico.Dices{
			dico.Dice{'R', 'I', 'F', 'O', 'B', 'X'},
			dico.Dice{'I', 'F', 'E', 'H', 'E', 'Y'},
			dico.Dice{'D', 'E', 'N', 'O', 'W', 'S'},
			dico.Dice{'U', 'T', 'O', 'K', 'N', 'D'},
			dico.Dice{'H', 'M', 'S', 'R', 'A', 'O'},
			dico.Dice{'L', 'U', 'P', 'E', 'T', 'S'},
			dico.Dice{'A', 'C', 'I', 'T', 'O', 'A'},
			dico.Dice{'Q', 'B', 'M', 'J', 'O', 'A'},
			dico.Dice{'E', 'H', 'I', 'S', 'P', 'N'},
			dico.Dice{'V', 'E', 'T', 'I', 'G', 'N'},
			dico.Dice{'G', 'L', 'E', 'N', 'Y', 'U'},
			dico.Dice{'B', 'A', 'L', 'I', 'Y', 'T'},
			dico.Dice{'E', 'Z', 'A', 'V', 'N', 'D'},
			dico.Dice{'R', 'A', 'L', 'E', 'S', 'C'},
			dico.Dice{'U', 'W', 'I', 'L', 'R', 'G'},
			dico.Dice{'P', 'A', 'C', 'E', 'M', 'D'},
		}
	*/

	lmn := dices.LetterMaxNumber()
	lom := dices.LetterOption()

	rp := dices.RemovePaternStruct(lmn, lom)

	dico.RemoveOfTxt(input, dices, rp, lmn)

	origin := tree.GenerateDicoFromTxt(input)

	if params[ui.EXPORT_ALL] == 1 {
		utils.Encode(lmn, filepath.Join(filepath.Dir(output), "lmn.json"))
		utils.Encode(lom, filepath.Join(filepath.Dir(output), "lom.json"))
		utils.Encode(rp, filepath.Join(filepath.Dir(output), "remove-patern.json"))
	}

	utils.Encode(origin, output)

	elapsed := time.Since(start)
	log.Printf("Took %s\n", elapsed)
}
