package main

import (
	"bytes"
	"encoding/json"
	"html/template"
)

type RequestBody struct {
	Mode string `json:"mode"`
	Raw  string `json:"raw"`
}

func (b RequestBody) RawString() template.HTML {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(b.Raw), "\t\t\t", "\t")
	if err != nil {
		return template.HTML(b.Raw)
	}
	return template.HTML(out.String())
}
