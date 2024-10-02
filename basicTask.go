package basicTask

func IsPalindrome(inStr string) bool {

	s := 0
	e := len(inStr)-1
	
	for ; s < e; {
		if inStr[s] != inStr[e] { return false }
		 s++; e--
	}
	return true
}
