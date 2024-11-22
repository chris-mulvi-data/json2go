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
