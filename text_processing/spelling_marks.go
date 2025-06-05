package text_processing

import (
	"go_reloaded/additional_functions"
	"regexp"
	"strings"
)

// CorrectPunctuation удаляет лишние пробелы перед знаками препинания и между ними.
// Использует регулярные выражения, определённые в additional_functions.
func CorrectPunctuation(text string) string {
	text = additional_functions.RemoveSpaceBeforePunct.ReplaceAllString(text, "$1")
	text = additional_functions.RemoveSpaceBetweenPuncts.ReplaceAllString(text, "$1$2")
	return text
}

// handleApostrophes обрабатывает кавычки-апострофы в тексте:
// - убирает лишние пробелы внутри одиночных кавычек;
// - корректно разделяет подряд идущие кавычки (например, ''foo'' → 'foo' 'foo').
func handleApostrophes(text string) string {
	// Удаляет пробелы внутри кавычек, например: ' hello ' → 'hello'
	processed := additional_functions.ApostropheContentWithSpaces.ReplaceAllStringFunc(text, func(match string) string {
		submatches := additional_functions.ApostropheContentWithSpaces.FindStringSubmatch(match)
		content := submatches[2] // Извлекаем содержимое внутри кавычек
		trimmedContent := strings.TrimSpace(content)

		// Проверка, заканчивается ли содержимое знаком препинания
		hasPunctuationAtEnd := false
		if len(trimmedContent) > 0 {
			lastChar := trimmedContent[len(trimmedContent)-1]
			hasPunctuationAtEnd = strings.ContainsRune(".,!?;:", rune(lastChar))
		}

		// Удаляет пробел перед пунктуацией, если она присутствует
		if hasPunctuationAtEnd {
			trimmedContent = additional_functions.RemoveSpaceBeforePunct.ReplaceAllString(trimmedContent, "$1")
		}

		return "'" + trimmedContent + "'"
	})

	// Обработка случая двойных кавычек подряд без пробела: 'foo''bar' → 'foo' 'bar'
	re1 := regexp.MustCompile(`'([^']*)''([^']*)'`)
	flagmatchstring, _ := regexp.MatchString(`'([^']*)''([^']*)'`, processed)
	for flagmatchstring {
		processed = re1.ReplaceAllString(processed, "'$1' '$2'")
		flagmatchstring, _ = regexp.MatchString(`'([^']*)''([^']*)'`, processed)
	}

	// Ещё раз применим шаблон, чтобы убедиться, что всё обработано
	text = re1.ReplaceAllString(processed, "'$1' '$2'")
	return text
}
