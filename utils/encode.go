package utils

import (
	"data_structure"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"ui"
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
				fmt.Print(err, '\n')
				ui.PrintHelp()
				os.Exit(0)
			}
		}

		encoder := json.NewEncoder(file)
		err = encoder.Encode(v)
		if err != nil {
			fmt.Print(err, '\n')
			ui.PrintHelp()
			os.Exit(0)
		}
	}

}

func EncodeNode(v *data_structure.Node, path ...string) {
	// si on souhaite exporter au format json
	if len(path) > 0 {
		out := filepath.Join(path...)

		file, err := os.Create(out)
		if err != nil {
			os.Mkdir(filepath.Dir(out), 777)
			file, err = os.Create(out)
			if err != nil {
				fmt.Print(err, '\n')
				ui.PrintHelp()
				os.Exit(0)
			}
		}

		encoder := json.NewEncoder(file)
		err = encoder.Encode(v.Convert())
		if err != nil {
			fmt.Print(err, '\n')
			ui.PrintHelp()
			os.Exit(0)
		}
	}

}
