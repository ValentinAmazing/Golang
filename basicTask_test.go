package basicTask

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// Arrange  тестовые данные
	testTable := []struct {
		word string
		expected bool
	}{
		{
			word: "",
			expected: true,
		},
		{
			word: "golang",
			expected: false,
		},
		{
			word: "level",
			expected: true,
		},
		{
			word: "abba",
			expected: true,
		},
		{
			word: "Race car",
			expected: true,
		},
	}

	// Act  тест 
	for _, testCase := range testTable {
		result := IsPalindrome(testCase.word)

	// Assert  Проверка результата
		if result != testCase.expected {
			t.Errorf( "Incorrect result. Expected '%t' for '%s'",
				testCase.expected, testCase.word)
		}		
	}

}
