package ui

import (
	"errors"
	"os"
	"regexp"
)

type Args string
type ArgsMap map[Args]int

const (
	EXPORT_ALL            Args = "--export-all"
	EXPORT_ALL_COMPRESSED Args = "-ea"
	SPECIAL_OUPUT         Args = "--out="
	DICE_FILE             Args = "--dice="
	DICE_FILE_COMPRESSED  Args = "-d="
	TARGET                Args = "t"
)

func GetParams() ArgsMap {

	res := make(ArgsMap)
	if len(os.Args) < 2 {
		panic(errors.New("No target"))
	}

	res[EXPORT_ALL] = -1
	res[SPECIAL_OUPUT] = -1

	for i, v := range os.Args[1:] {
		switch v {
		case string(EXPORT_ALL_COMPRESSED):
			res[EXPORT_ALL] = 1
		case string(EXPORT_ALL):
			res[EXPORT_ALL] = 1
		default:
			if ok, _ := regexp.Match(string(SPECIAL_OUPUT)+"*", []byte(v)); ok {
				res[SPECIAL_OUPUT] = i + 1
			} else {
				res[TARGET] = i + 1
			}
		}
	}

	return res
}
