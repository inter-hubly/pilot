package valueobject

import (
	"regexp"
	"strings"
)

type Slug struct {
	value string
}

func NewSlug(value string) *Slug {
	value = strings.ToLower(value)

	re := regexp.MustCompile(`[^a-z0-9\s-]`)
	value = re.ReplaceAllString(value, "")
	value = strings.Join(strings.Fields(value), "-")

	return &Slug{value: value}
}

func (s *Slug) Value() string {
	return s.value
}
