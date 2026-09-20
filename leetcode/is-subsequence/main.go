package main

func isSubsequence(s string, t string) bool {
	sptr := 0

	if len(s) == 0 {
		return true
	}

	for tptr := 0; tptr < len(t); tptr++ {
		if s[sptr] == t[tptr] {
			sptr++
		}
		if sptr == len(s) {
			return true
		}
	}
	return false
}

func main() {

}
