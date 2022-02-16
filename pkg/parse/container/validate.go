package container

const (
	openCurly  rune = '{'
	closeCurly rune = '}'
	openBox    rune = '['
	closeBox   rune = ']'
	openRound  rune = '('
	closeRound rune = ')'
)

var pairs = map[rune]rune{'{': '}', '[': ']', '(': ')'}

var btMap = map[rune]func(queue []rune, char rune) ([]rune, bool){
	openCurly:  openBracket,
	closeCurly: closeBracket,
	openBox:    openBracket,
	closeBox:   closeBracket,
	openRound:  openBracket,
	closeRound: closeBracket,
}

func BracketCheck(phrase string) (bool, error) {
	var queue []rune
	var checked bool
	for _, v := range phrase {
		if checkFunc, ok := btMap[v]; ok {
			queue, checked = checkFunc(queue, v)
			if !checked {
				return false, nil
			}
		}
	}
	return len(queue) == 0, nil
}

func openBracket(queue []rune, char rune) ([]rune, bool) {
	queue = append(queue, pairs[char])
	return queue, true
}

func closeBracket(queue []rune, char rune) ([]rune, bool) {
	if 0 < len(queue) && queue[len(queue)-1] == char {
		queue = queue[:len(queue)-1]
	} else {
		return queue, false
	}
	return queue, true
}
