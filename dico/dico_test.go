package dico

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestGetAllWord(t *testing.T) {

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
		for _, e := range res[0:10] {
			log.Println(e)
		}
		t.Log("Success")
	}

}
