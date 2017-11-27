package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/docopt/docopt-go"
)

var (
	table = map[string]string{
		"0": "no",
		"1": "pa",
		"2": "re",
		"3": "ci",
		"4": "vo",
		"5": "mu",
		"6": "xa",
		"7": "ze",
		"8": "bi",
		"9": "so",
	}

	retable = swap(table)
)

var (
	version = "[manual build]"
	usage   = "lojban " + version + `

Convert arabic numerals to lojban and backward.

Usage:
  lojban [options] <target>
  lojban -h | --help
  lojban --version

Options:
  -p --prefix <prefix>  Prefix before numerals.
  -h --help             Show this screen.
  --version             Show version.
`
)

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		panic(err)
	}

	var (
		prefix, _ = args["--prefix"].(string)
		target    = args["<target>"].(string)
	)

	if isNumber(target) {
		lojban := toLojban(target)

		fmt.Printf("%s%s\n", prefix, lojban)

		return
	}

	target = strings.TrimPrefix(target, prefix)

	numeric, err := fromLojban(target)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(numeric)
}

func fromLojban(target string) (string, error) {
	if len(target)%2 != 0 {
		return "", fmt.Errorf(
			"length of lojban string should be even, but got odd (%d)",
			len(target),
		)
	}

	result := ""
	for i := 0; i < len(target); i += 2 {
		word := string(target[i]) + string(target[i+1])

		conversion, ok := retable[word]
		if !ok {
			return "", fmt.Errorf("unexpected piece: %q", word)
		}

		result += conversion
	}

	return result, nil
}

func toLojban(target string) string {
	result := ""
	for _, sym := range target {
		result += table[string(sym)]
	}
	return result
}

func isNumber(target string) bool {
	for _, sym := range target {
		if sym < '0' || sym > '9' {
			return false
		}
	}

	return true
}

func swap(table map[string]string) map[string]string {
	swapped := map[string]string{}
	for key, value := range table {
		swapped[value] = key
	}
	return swapped
}
