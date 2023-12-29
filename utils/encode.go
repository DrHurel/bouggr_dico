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
			panic(err)
		}

		encoder := json.NewEncoder(file)
		encoder.Encode(v)
	}

}
