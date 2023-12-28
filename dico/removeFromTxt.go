package dico

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
)

func RemoveFromTxt(path string, dices Dices, rpPath string) {

	var bs []byte
	var text string
	var rp RemovePatern

	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	rpfile, err := os.ReadFile(rpPath)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(rpfile, &rp); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	buf := bytes.NewBuffer(bs)

	for scanner.Scan() {
		text = scanner.Text()

		length := len(text)

		if length <= 16 && length > 2 {
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
