package dico

import (
	"bufio"
	"data_structure"
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"
)

func TestGetAllWord(t *testing.T) {
	dicoTxt := "C:/Users/myyou/AppData/Local/git-r/boggle_dico/fr_dico_copy.txt"

	readFile, err := os.Open(dicoTxt)

	if err != nil {
		log.Fatalf("failed to open")
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	mapDico := make(map[string]bool, 407935)
	for fileScanner.Scan() {
		mapDico[fileScanner.Text()] = true
	}

	grid := "OVERFCEANAEBRUAS"
	dicofile, err := os.ReadFile("C:/Users/myyou/AppData/Local/git-r/boggle_dico/out/fr_dico.json")

	if err != nil {
		t.Errorf("Error opening file %s", err)
		t.Fail()
	}

	// json parse from file
	var d [2]interface{}

	err = json.Unmarshal(dicofile, &d)

	if err != nil {
		t.Errorf("Error decoding file")
		t.Fail()
	}
	if len(d) == 0 {
		t.Errorf("Error decoding file")
		t.Fail()
	}
	dc := data_structure.GenerateDicoFromTxt(dicoTxt, data_structure.LangueCodeMap([]string{"FR"}))
	dc2 := dc.Convert()
	start := time.Now()

	res := GetAllWord(grid, dc2.([]interface{}))
	elapse := time.Since(start)

	grid2 := [4][4]rune{
		{'O', 'V', 'E', 'R'},
		{'F', 'C', 'E', 'A'},
		{'N', 'A', 'E', 'B'},
		{'R', 'U', 'A', 'S'},
	}

	start2 := time.Now()
	res2 := AllWordInGrid(grid2, dc, 1<<9)
	elapse2 := time.Since(start2)
	if elapse > 500*time.Millisecond {
		t.Errorf("Expected <0.5s, got %s", elapse)
		t.Fail()
	} else {

		log.Println(elapse, elapse2)
	}
	if len(res) == 0 {
		t.Errorf("Expected >=1, got %d", len(res))
		t.Fail()
	} else {

		log.Println(len(res))
		log.Println(len(res2))
		if len(res) > 410000 {
			t.Errorf("Expected <410000, got %d", len(res))
			t.Fail()
		}
		count := 0
		mapRes1 := make(map[string]bool, len(res))
		for _, e := range res {
			mapRes1[e] = true
			if _, ok := mapDico[e]; !ok { //ensure that the word are also found in the dico

				t.Errorf("Expected %s, not in dico", e)

				count++
				//t.Errorf("Expected %s, not in dico", e)

			}
		}

		count2 := 0
		count3 := 0
		for _, e := range res2 {
			if mapRes1[e] == false { //ensure that the word are also found in the first result
				count3++
			}
			if _, ok := mapDico[e]; !ok { //ensure that the word are also found in the dico
				count2++
				t.Errorf("Expected %s, not in dico2", e)

			}

		}

		if count2 > 0 { // in first to verify that the second function is refenrentialy correct
			t.Errorf("Not ref %d, got %d out of %d (%d) (%d)", 0, count2, len(res2), len(res2)-count2, len(res))
			t.Fail()
		}

		if count > 0 {

			t.Errorf("Not %d, got %d out of %d (%d) (%d)", 0, count, len(res), len(res)-count, len(res2))
			t.Fail()
		}

		if count3 > 0 {

			t.Errorf("Not 3 %d, got %d out of %d (%d) (%d)", 0, count3, len(res2), len(res2)-count3, len(res))
			t.Fail()
		}

	}

}
