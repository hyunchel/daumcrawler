package daumcrawler

import (
	"strings"
	"testing"
)

func TestConvertCamelCase(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"some keyword", "someKeyword"},
		{"another keyword that has more spaces", "anotherKeywordThatHasMoreSpaces"},
		{"", ""},
		{"keyword", "keyword"},
		{"keyWord", "keyWord"},
		{"simPle keyWord", "simPleKeyWord"},
	}
	for _, c := range cases {
		got := convertToCamelCase(c.in)
		if got != c.want {
			t.Errorf("Expected %q to yield %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMakeFilename(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"some keyword", "-someKeyword"},
		{"another keyword that has more spaces", "-anotherKeywordThatHasMoreSpaces"},
		{"", "-"},
		{"keyword", "-keyword"},
	}
	for _, c := range cases {
		got := makeFilename(c.in)
		if strings.HasSuffix(got, c.want) == false {
			t.Errorf("Expected %q in %q, want %q", c.in, got, c.want)
		}
	}
}
