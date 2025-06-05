package text_processing

import (
	"go_reloaded/additional_functions"
	"math/big"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Transform описывает трансформацию: функция и количество слов, к которым она применяется
type Transform struct {
	fn    func(string) string
	count int
}

// reverseSlice переворачивает срез строк — используется для обработки тэгов справа налево
func reverseSlice(slice []string) []string {
	reversed := make([]string, len(slice))
	for i, j := 0, len(slice)-1; i < len(slice); i, j = i+1, j-1 {
		reversed[i] = slice[j]
	}
	return reversed
}

// ProcessTags применяет трансформации по тегам: (up), (low), (cap), (hex), (bin)
// Теги могут иметь форму (tag,count) — например: (up,3)
func ProcessTags(tokens []string) []string {
	// Переворачиваем токены для обратной обработки
	reversed := reverseSlice(tokens)
	transformed := []string{}
	activeTransforms := []Transform{}

	// Проверка: является ли токен тегом вида (xxx)
	isTag := func(token string) bool {
		return strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")")
	}

	// Проверка: является ли токен словом (не тег и не пунктуация)
	isWord := func(token string) bool {
		return !isTag(token) && !additional_functions.IsPunctuation(token)
	}

	// Набор поддерживаемых трансформаций
	transformations := map[string]func(string) string{
		"up":  strings.ToUpper,
		"low": strings.ToLower,
		"cap": func(s string) string {
			if s == "" {
				return ""
			}

			// Обработка слова, начинающегося с кавычки
			if strings.HasPrefix(s, "'") || strings.HasPrefix(s, "\"") {
				trimmed := s[1:]
				r, size := utf8.DecodeRuneInString(trimmed)
				return string(s[0]) + string(unicode.ToUpper(r)) + strings.ToLower(trimmed[size:])
			}

			// Обычная капитализация: первая буква заглавная, остальные строчные
			r, size := utf8.DecodeRuneInString(s)
			return string(unicode.ToUpper(r)) + strings.ToLower(s[size:])
		},
		"hex": func(s string) string {
			// Преобразование из HEX в десятичное, если строка — валидный hex
			if additional_functions.IsHex(s) {
				n := new(big.Int)
				if _, success := n.SetString(s, 16); success {
					return n.String()
				}
			}
			return s
		},
		"bin": func(s string) string {
			// Преобразование из BIN в десятичное, если строка — валидный бинарный
			if additional_functions.IsBinary(s) {
				n := new(big.Int)
				if _, success := n.SetString(s, 2); success {
					return n.String()
				}
			}
			return s
		},
	}

	// Основной цикл обработки токенов
	for _, token := range reversed {
		if isTag(token) {
			// Разбираем тег и его параметры
			tagContent := token[1 : len(token)-1]
			parts := strings.Split(tagContent, ",")
			transformation := strings.ToLower(strings.TrimSpace(parts[0]))

			// Если тег неизвестен — сохраняем как есть
			if _, ok := transformations[transformation]; !ok {
				transformed = append(transformed, token)
				continue
			}

			// Обработка второго параметра (кол-во слов)
			count := 1
			if len(parts) > 1 {
				countStr := strings.TrimSpace(parts[1])
				parsedCount, err := strconv.Atoi(countStr)
				if err != nil || parsedCount <= 0 {
					continue // пропускаем тег с некорректным числом
				}
				count = parsedCount
			}

			// Добавляем трансформацию в стек активных
			if transformFn, ok := transformations[transformation]; ok {
				activeTransforms = append(activeTransforms, Transform{fn: transformFn, count: count})
			}
		} else {
			// Если это слово, применяем активные трансформации
			if isWord(token) && len(activeTransforms) > 0 {
				for i := len(activeTransforms) - 1; i >= 0; i-- {
					if activeTransforms[i].count > 0 {
						token = activeTransforms[i].fn(token)
						activeTransforms[i].count--
					}
				}
			}
			transformed = append(transformed, token)

			// Убираем трансформации, у которых счётчик = 0
			var newActive []Transform
			for _, t := range activeTransforms {
				if t.count > 0 {
					newActive = append(newActive, t)
				}
			}
			activeTransforms = newActive
		}
	}

	// Возвращаем в исходном порядке
	return reverseSlice(transformed)
}
