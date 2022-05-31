func canConstruct(ransomNote string, magazine string) bool {
	var found bool = false
	for _, b := range ransomNote {
		magazine, found = find(b, magazine)
		if !found {
			return false
		}
	}
	return found
}

func find(r rune, s string) (string, bool) {
	for i, a := range s {
		if a == r {
			r := s[:i]
			r2 := s[i+1:]
			return r + r2, true
		}
	}
	return s, false
}
