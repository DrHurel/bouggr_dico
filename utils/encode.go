package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Encode(v any, path ...string) {
	// si on souhaite exporter au format json
	if len(path) > 0 {
		out := filepath.Join(path...)

		file, err := os.Create(out)
		if err != nil {
			os.Mkdir(filepath.Dir(out), 777)
			file, err = os.Create(out)
			if err != nil {
				panic(err)
			}
		}

		encoder := json.NewEncoder(file)
		err = encoder.Encode(v)
		if err != nil {
			panic(err)
		}
	}

}
