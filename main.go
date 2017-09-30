// gotmpl process template from STDIN to STDOUT using text/template.
//
// Data available in template:
//  .Env  - map[string]string with environment variables
//
// Functions available in template:
//  json  - encode as JSON string in double-quotes
//  trimq - remove surrounding double-quotes, if any
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"
)

var (
	app = strings.TrimSuffix(path.Base(os.Args[0]), ".test")
	ver = "v?"
	cfg struct {
		version bool
	}
	funcMap = template.FuncMap{
		"json":  escapeJSON,
		"trimq": trimQ,
	}
	data struct {
		Env map[string]string
	}
)

func init() {
	flag.BoolVar(&cfg.version, "version", false, "print version")
	data.Env = env()
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
		os.Exit(2)
	}

	if cfg.version {
		fmt.Println(app, ver, runtime.Version())
		os.Exit(0)
	}

	err := tmpl(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func tmpl(tmpl io.Reader, text io.Writer) error {
	buf, err := ioutil.ReadAll(tmpl)
	if err != nil {
		return err
	}
	t, err := template.New("main").Funcs(funcMap).Parse(string(buf))
	if err != nil {
		return err
	}
	return t.Execute(text, data)
}

func env() map[string]string {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		env[kv[0]] = kv[1]
	}
	return env
}

func escapeJSON(v interface{}) (string, error) {
	var b bytes.Buffer
	e := json.NewEncoder(&b)
	e.SetEscapeHTML(false)
	err := e.Encode(v)
	return strings.TrimSpace(b.String()), err
}

func trimQ(s string) string {
	s = strings.TrimPrefix(s, `"`)
	s = strings.TrimSuffix(s, `"`)
	return s
}
