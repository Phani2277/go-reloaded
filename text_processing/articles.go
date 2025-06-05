package text_processing

import (
	"go_reloaded/additional_functions"
	"regexp"
	"strings"
	"unicode"
)

// CorrectArticles корректирует неопределённые артикли "a" и "an" в переданном тексте.
// Учитываются начальные звуки следующих слов (гласные, "молчаливое h", и исключения).
func CorrectArticles(text string) string {
	// Заменяем переносы строк на специальный маркер, чтобы сохранить структуру
	text = strings.ReplaceAll(text, "\n", "NEWLINE_MARKER")
	// Добавляем пробелы вокруг маркера, чтобы избежать склеивания с соседними словами
	text = strings.ReplaceAll(text, "NEWLINE_MARKER", " NEWLINE_MARKER ")
	words := strings.Fields(text) // Разбиваем текст на слова

	if len(words) == 0 {
		// Если текст пустой, возвращаем его с восстановленными переносами
		return strings.ReplaceAll(text, "NEWLINE_MARKER", "\n")
	}

	result := make([]string, len(words)) // Слайс для результата

	// Проходим по каждому слову
	for i := 0; i < len(words); i++ {
		word := words[i]
		result[i] = word // По умолчанию оставляем слово как есть

		lowerWord := strings.ToLower(word)
		// Пропускаем слова, не являющиеся артиклями
		if lowerWord != "a" && lowerWord != "an" {
			continue
		}

		// Если это последнее слово — не с чем сравнивать, пропускаем
		if i+1 >= len(words) {
			continue
		}

		nextWord := words[i+1]
		if len(nextWord) == 0 {
			continue
		}

		// Пропускаем, если следующее слово — артикль, пунктуация или союз
		if additional_functions.IsArticle(nextWord) {
			continue
		}
		lowerNextWord := strings.ToLower(nextWord)
		if additional_functions.IsPunctuation(nextWord) || additional_functions.IsConjunction(lowerNextWord) {
			continue
		}

		shouldBeAn := false // Флаг, нужен ли артикль "an"
		runes := []rune(strings.ToLower(nextWord))

		// Определяем первую значимую букву (игнорируя знаки препинания)
		firstChar := '0'
		if !unicode.IsLetter(runes[0]) {
			firstChar = runes[1]
		} else {
			firstChar = runes[0]
		}

		// Обработка исключений с "молчаливым h"
		if firstChar == 'h' {
			silentHWords := []string{
				"hour", "honor", "honour", "honest", "heir", "herb",
			}
			cleanNextWord := strings.Trim(lowerNextWord, ".,!?;:'")
			silentHWordsFlag := false
			for _, silentHWord := range silentHWords {
				if strings.HasPrefix(cleanNextWord, silentHWord) {
					silentHWordsFlag = true
					break
				}
			}
			if silentHWordsFlag {
				shouldBeAn = true
			}
		} else if strings.ContainsRune("aeiou", firstChar) {
			// Если слово начинается с гласной
			shouldBeAn = true
		}

		// Исключения для "u" — например, "university", "european"
		if firstChar == 'u' {
			if strings.HasPrefix(lowerNextWord, "uni") || strings.HasPrefix(lowerNextWord, "eu") {
				shouldBeAn = false
			}
		}

		// Корректируем артикль в зависимости от анализа
		if shouldBeAn {
			if word == "a" {
				result[i] = "an"
			} else if word == "A" && !unicode.IsLower(rune(nextWord[0])) && !unicode.IsLower(rune(nextWord[1])) {
				result[i] = "AN"
			} else if word == "A" {
				result[i] = "An"
			}
		} else {
			if word == "an" {
				result[i] = "a"
			} else if word == "An" || word == "AN" {
				result[i] = "A"
			}
		}
	}

	// Восстанавливаем структуру текста
	out := strings.Join(result, " ")
	out = strings.ReplaceAll(out, "NEWLINE_MARKER", "\n")
	// Удаляем лишние пробелы перед и после перевода строки
	out = regexp.MustCompile(`[ \t]+\n`).ReplaceAllString(out, "\n")
	out = regexp.MustCompile(`\n[ \t]+`).ReplaceAllString(out, "\n")

	return out
}
