package dico

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"testing"
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

	t.Log(d[0])
	res := GetAllWord(grid, d)
	if len(res) == 0 {
		t.Errorf("Expected >=1, got %d", len(res))
		t.Fail()
	} else {

		log.Println(len(res))

		if len(res) > 410000 {
			t.Errorf("Expected <410000, got %d", len(res))
			t.Fail()
		}
		count := 0
		for _, e := range res {

			if _, ok := mapDico[e]; !ok {
				count++
				t.Errorf("Expected %s, not in dico", e)

			}
		}

		if count > 0 {

			t.Errorf("Not %d, got %d out of %d (%d)", 0, count, len(res), len(res)-count)
			t.Fail()
		}

	}

}
