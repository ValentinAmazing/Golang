package basicTask

func IsPalindrome(str string) bool {

	s := 0
	e := len(str)-1
	
	for ; s < e; {
		if str[s] != str[e] { return false }
		 s++; e--
	}
	return true
}
