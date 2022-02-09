package container

func IsValid(s string) bool {
	var brakets []byte
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			brakets = append(brakets, byte(v))
		case ')', ']', '}':
			if len(brakets) == 0 {
				return false
			}
			if v == ')' && brakets[len(brakets)-1] == '(' ||
				v == ']' && brakets[len(brakets)-1] == '[' ||
				v == '}' && brakets[len(brakets)-1] == '{' {
				brakets = brakets[:len(brakets)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}
	if len(brakets) > 0 {
		return false
	}
	return true
}
