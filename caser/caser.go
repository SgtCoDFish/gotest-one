package caser

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Dutch(s string) string {
	return cases.Title(language.Dutch).String(s)
}

func Turkish(s string) string {
	return cases.Title(language.Turkish).String(s)
}
