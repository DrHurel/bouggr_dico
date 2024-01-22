package main

import (
	"data_structure"
	"dico"
	"fmt"
	"stat"

	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"ui"
	"utils"
)

func main() {

	var wg sync.WaitGroup

	var output string
	var lmn dico.IterationCount
	var lom dico.LetterOptionMap

	// Get Params
	params, option := ui.GetParams()
	if option {
		ui.PrintHelp()
		return
	}

	input := os.Args[1]

	output = ui.GetOutput(params, output)

	fmt.Print("Generate\n")
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

	if params[ui.RULES_FILES] != -1 {
		dices = dico.GetDicesFromTxt(strings.Split(os.Args[params[ui.RULES_FILES]], "=")[1])
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

	wg.Add(2)
	go func() {
		lmn = dices.LetterMaxNumber()
		defer wg.Done()
	}()

	go func() {
		lom = dices.LetterOption()
		defer wg.Done()
	}()

	wg.Wait()
	rp := dices.RemovePaternStruct(lmn, lom)
	langs := data_structure.LangueCodeMap(ui.GetLangList(params))

	noParse := data_structure.GenerateDicoFromTxt(input, langs)
	if params[ui.NO_PARSE] == -1 {
		dico.RemoveOfTxt(input, dices, rp, lmn)
	}

	origin := data_structure.GenerateDicoFromTxt(input, langs)
	d := []rune{}
	for _, dice := range dices {
		d = append(d, dice...)
	}

	if params[ui.EXPORT_ALL] == 1 {
		wg.Add(3)
		go func() {
			utils.Encode(lmn, filepath.Join(filepath.Dir(output), "lmn.json"))
			defer wg.Done()
		}()
		go func() {
			utils.Encode(lom, filepath.Join(filepath.Dir(output), "lom.json"))
			defer wg.Done()
		}()
		go func() {
			utils.Encode(rp, filepath.Join(filepath.Dir(output), "remove-patern.json"))
			defer wg.Done()
		}()
	}

	utils.Encode(origin, output)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)

	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond()) * elapsed.Nanoseconds()))

	grid := dices.Roll(r)

	fmt.Print("\n")
	utils.PrintGrid(grid)

	fmt.Print("\n")
	fmt.Print(origin.CheckWord("AAS", 1), origin.CheckWord("ABAISSABLES", 1), origin.CheckWord("ABAISSjsoidngiodj,vABLES", 1), origin.CheckWord("ZINZINULENT", 1), origin.CheckWord("IAERBSRUGANOS", 1))
	start = time.Now()
	fmt.Println("Start")
	allw := dico.AllWordInGrid(grid, origin)
	fmt.Println("Stop")
	elapsed = time.Since(start)

	for _, e := range allw {
		fmt.Println(e)
	}
	fmt.Printf("\nTook %s\n", elapsed)

	if params[ui.FORCE] != 1 {
		stat.Display(grid, dices, r, start, origin, noParse)
	}

	os.Exit(0)
}
