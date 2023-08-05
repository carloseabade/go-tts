package main

import (
	"fmt"
	"os"
	//"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	c := map[string]rune{
		"period":      '.',
		"exclamation": '!',
		"question":    '?',
	}

	s := make([]string, 1)
	lastCut := 0
	lastIndex := 0
	for i, v := range string(dat) {
    if v == ' ' {
      for k := range c {
        if rune(string(dat)[lastIndex]) == c[k] {
          s = appendToS(s, string(dat)[lastCut:i+1], &i, &lastCut)
          break;
        }
      }
    }
    lastIndex = i
	}

  for i := range s {
    fmt.Printf("%s\n", s[i])
  }
}

func appendToS(s []string, str string, i, lC *int) []string {
  s = append(s, str)
  *i = *i + 1
  *lC = *i
  return s
}
