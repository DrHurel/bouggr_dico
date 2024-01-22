package ui

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Args string
type ArgsMap map[Args]int
type conversionMap map[Args]Args

const (
	NO_ARGS                      = -1
	EXPORT_ALL            Args   = "--export-all"
	EXPORT_ALL_COMPRESSED Args   = "-ea"
	SPECIAL_OUPUT         Args   = "--out="
	RULES_FILES           Args   = "--rules="
	LANGUE_LIST           Args   = "--lang="
	NO_PARSE              Args   = "--no-parse"
	NO_PARSE_COMPRESSED   Args   = "-np"
	FORCE                 Args   = "--force"
	FORCE_COMPRESSED      Args   = "-F"
	HELP                  Args   = "--help"
	STANDARD_OUTPUT       string = "./out/dico.json"
)

func GetParams() (ArgsMap, bool) {

	if len(os.Args) < 2 {
		PrintHelp()
		os.Exit(0)
	}

	if os.Args[1] == string(HELP) {
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
		EXPORT_ALL:    NO_ARGS,
		SPECIAL_OUPUT: NO_ARGS,
		RULES_FILES:   NO_ARGS,
		NO_PARSE:      NO_ARGS,
		FORCE:         NO_ARGS,
	}

	for i, v := range os.Args[1:] {

		if ok, _ := regexp.Match(string(SPECIAL_OUPUT)+"*", []byte(v)); ok {
			res[SPECIAL_OUPUT] = i + 1
		}
		if ok, _ := regexp.Match(string(RULES_FILES)+"*", []byte(v)); ok {
			res[RULES_FILES] = i + 1
		}

		if ok, _ := regexp.Match(string(LANGUE_LIST)+"*", []byte(v)); ok {
			res[LANGUE_LIST] = i + 1
		}

		if _, found := conversion[Args(v)]; found {
			res[conversion[Args(v)]] = 1
		}

	}

	return res, false
}

func GetOutput(params ArgsMap, output string) string {
	if params[SPECIAL_OUPUT] != NO_ARGS {
		output = strings.Split(os.Args[params[SPECIAL_OUPUT]], "=")[1]
	} else {
		output = STANDARD_OUTPUT
	}
	return output
}

func GetLangList(params ArgsMap) []string {

	if params[LANGUE_LIST] != NO_ARGS {
		res := strings.Split(strings.Split(os.Args[params[LANGUE_LIST]], "=")[1], ",")
		for _, l := range res {
			fmt.Print(l, " ")
		}
		return res
	}
	return []string{"FR"}
}
