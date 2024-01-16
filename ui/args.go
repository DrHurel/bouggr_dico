package ui

import (
	"errors"
	"os"
	"regexp"
)

type Args string
type ArgsMap map[Args]int
type conversionMap map[Args]Args

const (
	EXPORT_ALL            Args = "--export-all"
	EXPORT_ALL_COMPRESSED Args = "-ea"
	SPECIAL_OUPUT         Args = "--out="
	RULES_FILES           Args = "--rules="
	NO_PARSE              Args = "--no-parse"
	NO_PARSE_COMPRESSED   Args = "-np"
	FORCE                 Args = "--force"
	FORCE_COMPRESSED      Args = "-F"
	HELP                  Args = "--help"
)

func GetParams() (ArgsMap, bool) {

	if len(os.Args) < 2 {
		panic(errors.New("No target"))
	}

	if os.Args[0] == string(HELP) {
		return nil, true
	}

	conversion := conversionMap{
		EXPORT_ALL:            EXPORT_ALL,
		EXPORT_ALL_COMPRESSED: EXPORT_ALL,
		NO_PARSE:              NO_PARSE,
		NO_PARSE_COMPRESSED:   NO_PARSE,
		FORCE:                 FORCE,
		FORCE_COMPRESSED:      FORCE,
	}

	res := ArgsMap{
		EXPORT_ALL:    -1,
		SPECIAL_OUPUT: -1,
		RULES_FILES:   -1,
		NO_PARSE:      -1,
		FORCE:         -1,
	}

	for i, v := range os.Args[1:] {

		if ok, _ := regexp.Match(string(SPECIAL_OUPUT)+"*", []byte(v)); ok {
			res[SPECIAL_OUPUT] = i + 1
		}
		if ok, _ := regexp.Match(string(RULES_FILES)+"*", []byte(v)); ok {
			res[RULES_FILES] = i + 1
		}
		if _, found := conversion[Args(v)]; found {
			res[conversion[Args(v)]] = 1
		}

	}

	return res, false
}
