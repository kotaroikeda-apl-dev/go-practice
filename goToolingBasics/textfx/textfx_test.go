package textfx

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNormalizeSpaces(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   string
		want string
	}{
		{in: "  Go   is   fun  ", want: "Go is fun"},
		{in: "Hello,\tworld\n", want: "Hello, world"},
		{in: "", want: ""},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()
			got := NormalizeSpaces(tc.in)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("NormalizeSpaces mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRedactDigits(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   string
		want string
	}{
		{in: "order-2026-01-01", want: "order-xxxx-xx-xx"},
		{in: "電話: 090-1234-5678", want: "電話: xxx-xxxx-xxxx"},
		{in: "", want: ""},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()
			got := RedactDigits(tc.in, 'x')
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("RedactDigits mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTruncateRunes(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		in       string
		max      int
		ellipsis string
		want     string
	}{
		{name: "no-trunc", in: "Go", max: 2, ellipsis: "…", want: "Go"},
		{name: "truncate-ascii", in: "abcdef", max: 4, ellipsis: "…", want: "abc…"},
		{name: "truncate-jp", in: "こんにちは世界", max: 5, ellipsis: "…", want: "こんにち…"},
		{name: "max-too-small", in: "abcdef", max: 1, ellipsis: "…", want: "…"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := TruncateRunes(tc.in, tc.max, tc.ellipsis)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("TruncateRunes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSlugKebab(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   string
		want string
	}{
		{in: "  How to Write Go Code  ", want: "how-to-write-go-code"},
		{in: "Go Practice 2026!", want: "go-practice-2026"},
		{in: "Hello, 世界", want: "hello-世界"},
		{in: "---", want: ""},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()
			got := SlugKebab(tc.in)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("SlugKebab mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestApplyAll(t *testing.T) {
	t.Parallel()

	steps := []Step{
		NormalizeSpaces,
		func(s string) string { return RedactDigits(s, 'x') },
		func(s string) string { return TruncateRunes(s, 12, "…") },
		SlugKebab,
	}

	got := ApplyAll("  Order 2026-01-01  ", steps...)
	want := "order-xxxx"

	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("ApplyAll mismatch (-want +got):\n%s", diff)
	}
}

