package container

type bracketType int

const (
	openBracket bracketType = iota
	closeBracket
	notABracket
)

var pairs = map[rune]rune{'{': '}', '[': ']', '(': ')'}

func BracketCheck(phrase string) (bool, error) {
	var queue []rune
	for _, v := range phrase {
		switch getBracketType(v) {
		case openBracket:
			queue = append(queue, pairs[v])
		case closeBracket:
			if 0 < len(queue) && queue[len(queue)-1] == v {
				queue = queue[:len(queue)-1]
			} else {
				return false, nil
			}
		}
	}
	return len(queue) == 0, nil
}

func getBracketType(char rune) bracketType {
	for k, v := range pairs {
		switch char {
		case k:
			return openBracket
		case v:
			return closeBracket
		}
	}
	return notABracket
}
