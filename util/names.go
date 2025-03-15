package util

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// ToSlug normalize to slug
func ToSlug(name string, needUnderline bool) string {
	name = strings.ToLower(name)

	var sb strings.Builder
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == ' ' {
			sb.WriteRune(r)
		}
	}
	re := regexp.MustCompile(`[^a-z0-9-]`)
	name = re.ReplaceAllString(name, "")

	name = sb.String()
	var myChar string
	if needUnderline {
		myChar = "_"
	} else {
		myChar = "-"
	}

	name = strings.ReplaceAll(name, " ", myChar)
	name = strings.ReplaceAll(name, fmt.Sprintf("%s%s", myChar, myChar), myChar)
	name = strings.Trim(name, myChar)

	return name
}
