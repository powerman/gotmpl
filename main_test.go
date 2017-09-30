package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestTMPL(t *testing.T) {
	data.Env["a"] = `a"b` + "\n" + `c<d>e'f`
	var (
		template = "…{{ .Env.a |json|trimq}}…"
		exp      = `…a\"b\nc<d>e'f…`
		b        bytes.Buffer
	)

	err := tmpl(strings.NewReader(template), &b)
	if err != nil {
		t.Fatal(err)
	}
	res := b.String()
	if res != exp {
		t.Errorf(" = %v, want %v", res, exp)
	}
}

func TestOption(t *testing.T) {
	delete(data.Env, "a")
	var (
		template = "…{{ .Env.a }}…"
		exp      = `……`
		b        bytes.Buffer
	)

	err := tmpl(strings.NewReader(template), &b)
	if err != nil {
		t.Fatal(err)
	}
	res := b.String()
	if res != exp {
		t.Errorf(" = %v, want %v", res, exp)
	}
}
