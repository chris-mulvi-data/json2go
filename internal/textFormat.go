package internal

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PropertyToField(propertyName string) (res string, err error) {
	var formattedPropertyName string

	if strings.Contains(propertyName, "_") {
		parts := strings.Split(propertyName, "_")

		for _, part := range parts {
			formattedPropertyName += cases.Title(language.English).String(part)
		}
	}
	return formattedPropertyName, nil

}

// PadRightToSize addes white space to the right of a string
// until the size of the string reaches the deignated character count
func PadRightToSize(word *string, size int) {
	if len(*word) >= size {
		return
	}
	for {
		if len(*word) == size {
			break
		}
		*word += " "
	}
}

func AlignColumns(rows []string) []string {

	var charCounts map[int]int
	charCounts = make(map[int]int)

	for i := range rows {
		for j, word := range strings.Split(rows[i], " ") {
			value, isInMap := charCounts[j]
			if isInMap && value > len(word) {
				continue
			}
			charCounts[j] = len(word)
		}
	}

	for i := range rows {
		words := strings.Split(rows[i], " ")
		for j := range words {
			if j == len(words)-1 {
				break
			}
			PadRightToSize(&words[j], charCounts[j])
		}
		rows[i] = strings.Join(words, " ")
	}
	return rows
}
