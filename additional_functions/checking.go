package additional_functions

import (
	"regexp"
	"strings"
)

// IsPunctuation проверяет, является ли переданный токен пунктуацией.
// Использует регулярное выражение Checking_Punctuation.
func IsPunctuation(token string) bool {
	return Checking_Punctuation.MatchString(token)
}

// IsWord проверяет, является ли токен словом (не пунктуация и не заключён в круглые скобки).
func IsWord(token string) bool {
	return !IsPunctuation(token) && !strings.HasPrefix(token, "(") && !strings.HasSuffix(token, ")")
}

// IsHex проверяет, является ли строка допустимым шестнадцатеричным числом.
// Использует регулярное выражение IsHexCheck.
func IsHex(s string) bool {
	match, _ := regexp.MatchString(IsHexCheck, s)
	return match
}

// IsBinary проверяет, является ли строка допустимым двоичным числом (содержит только 0 и 1).
func IsBinary(s string) bool {
	for _, c := range s {
		if c != '0' && c != '1' {
			return false
		}
	}
	return true
}

// IsArticle проверяет, является ли слово неопределённым артиклем "a" или "an" (в любом регистре).
func IsArticle(word string) bool {
	return word == "a" || word == "an" || word == "A" || word == "An"
}

// IsConjunction проверяет, является ли слово союзом ("and", "or", "but", "nor").
func IsConjunction(word string) bool {
	conjunctions := map[string]bool{
		"and": true,
		"or":  true,
		"but": true,
		"nor": true,
	}
	return conjunctions[word]
}
