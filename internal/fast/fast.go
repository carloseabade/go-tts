package fast

var (
	cs = map[string]rune{
		"period":      '.',
		"exclamation": '!',
		"question":    '?',
	}
)

func SplitText(text string, c chan string) {
	lastIndex := 0
	for i, v := range string(text) {
    if v == ' ' {
      for k := range cs {
        if rune(string(text)[lastIndex]) == cs[k] {
					c <- string(text)[lastIndex:i+1]
          break;
        }
      }
    }
    lastIndex = i
	}
}

func appendToS(s []string, str string, i, lC *int) []string {
  s = append(s, str)
  *i = *i + 1
  *lC = *i
  return s
}
