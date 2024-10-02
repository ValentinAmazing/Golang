package basicTask

import "strings"

func IsPalindrome(inStr string) bool {

	inStr = strings.ReplaceAll(inStr, " ", "")
	inStr = strings.ToLower(inStr)
	s := 0
	e := len(inStr)-1
	
	for ; s < e; {
		if inStr[s] != inStr[e] { return false }
		 s++; e--
	}
	return true
}
