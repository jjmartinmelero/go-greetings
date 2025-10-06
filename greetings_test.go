package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Juan Jesus"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Juan Jesus")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(Juan Jesus) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestEmptyHello(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
