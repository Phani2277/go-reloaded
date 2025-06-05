package text_processing

import (
	"bytes"
	"go_reloaded/additional_functions"
	"strings"
)

// joinTokens соединяет токены обратно в текст, расставляя пробелы и корректируя пунктуацию.
func joinTokens(tokens []string) string {
	if len(tokens) == 0 {
		return ""
	}

	var buf bytes.Buffer
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		
		// Сохраняем перенос строки как есть
		if token == "\n" {
			buf.WriteString("\n")
			continue
		}

		// Первый токен или токен после новой строки — без пробела
		if i == 0 || tokens[i-1] == "\n" {
			buf.WriteString(token)
			continue
		}

		// Если это пунктуация, пишем без пробела перед ней
		if additional_functions.IsPunctuation(token) {
			buf.WriteString(token)
			continue
		}

		// Обычное слово — добавляем пробел перед ним
		buf.WriteString(" " + token)
	}

	result := buf.String()
	
	// Удаляем лишние пробелы перед пунктуацией
	result = additional_functions.RemoveSpaceBeforePunct.ReplaceAllString(result, "$1")
	// Удаляем пробелы между двумя пунктуационными знаками
	result = additional_functions.RemoveSpaceBetweenPuncts.ReplaceAllString(result, "$1$2")

	return result
}

// tokenize разбивает текст на токены, включая знаки препинания и переносы строк.
func tokenize(text string) []string {
	// Используем регулярное выражение из additional_functions для токенизации
	tokens := additional_functions.RegToken.FindAllString(text, -1)

	// Заменяем токены, содержащие перенос строки, на символ '\n'
	for i, token := range tokens {
		if strings.Contains(token, "\n") {
			tokens[i] = "\n"
		}
	}
	return tokens
}

// ProcessText выполняет все этапы обработки текста: токенизация, трансформация, корректировка пунктуации,
// обработка апострофов и исправление артиклей.
func ProcessText(text string) string {
	tokens := tokenize(text)                  // Токенизация
	transformedTokens := ProcessTags(tokens)  // Обработка пользовательских тегов (реализация отдельно)
	result := joinTokens(transformedTokens)   // Объединение токенов в строку
	result = CorrectPunctuation(result)       // Корректировка пунктуации
	result = handleApostrophes(result)        // Обработка апострофов
	result = CorrectArticles(result)          // Исправление артиклей ("a"/"an")
	return result
}
