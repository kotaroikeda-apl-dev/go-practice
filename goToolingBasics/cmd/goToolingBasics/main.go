package main

import (
	"flag"
	"fmt"

	"example.com/go-practice/goToolingBasics/textfx"
)

func main() {
	text := flag.String("text", "  How to Write Go Code 2026!  ", "input text")
	redact := flag.Bool("redact-digits", true, "mask ASCII digits [0-9]")
	max := flag.Int("max-runes", 32, "truncate output to at most N runes (0 disables)")
	slug := flag.Bool("slug", true, "convert to kebab-case slug")
	flag.Parse()

	steps := []textfx.Step{textfx.NormalizeSpaces}

	if *redact {
		steps = append(steps, func(s string) string { return textfx.RedactDigits(s, 'x') })
	}
	if *max > 0 {
		steps = append(steps, func(s string) string { return textfx.TruncateRunes(s, *max, "…") })
	}
	if *slug {
		steps = append(steps, textfx.SlugKebab)
	}

	out := textfx.ApplyAll(*text, steps...)
	fmt.Println(out)
}
