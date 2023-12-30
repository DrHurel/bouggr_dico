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
