// Package textfx provides small, composable text transformations.
//
// This package intentionally stays tiny and opinionated. It exists to demonstrate:
// - a module that contains multiple packages
// - importing your own package from a command under cmd/
// - writing tests (and managing a small external dependency for tests)
package textfx

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Step is a single text transformation.
type Step func(string) string

// ApplyAll applies steps from left to right.
func ApplyAll(s string, steps ...Step) string {
	for _, step := range steps {
		s = step(s)
	}
	return s
}

// NormalizeSpaces collapses all whitespace runs into a single ASCII space and trims ends.
func NormalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// RedactDigits replaces ASCII digits [0-9] with mask.
func RedactDigits(s string, mask rune) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(mask)
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// TruncateRunes returns s truncated to at most max runes.
// If truncation happens, ellipsis is appended (also counted toward max).
func TruncateRunes(s string, max int, ellipsis string) string {
	if max <= 0 {
		return ""
	}

	if utf8.RuneCountInString(s) <= max {
		return s
	}

	ell := []rune(ellipsis)
	keep := max - len(ell)
	if keep <= 0 {
		return string(ell[:max])
	}

	r := []rune(s)
	return string(r[:keep]) + ellipsis
}

// SlugKebab converts s into a simple "kebab-case" slug.
// Rules (deliberately small & deterministic):
// - letters/digits become lower-case
// - whitespace and punctuation become separators ('-')
// - separators are collapsed and trimmed
func SlugKebab(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	needSep := false
	wrote := false

	for _, r := range s {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			if wrote && needSep {
				b.WriteByte('-')
			}
			b.WriteRune(unicode.ToLower(r))
			wrote = true
			needSep = false
		default:
			if wrote {
				needSep = true
			}
		}
	}
	return b.String()
}

// Wrap returns a Step that surrounds the string with prefix and suffix.
func Wrap(prefix, suffix string) Step {
	return func(s string) string {
		return prefix + s + suffix
	}
}
