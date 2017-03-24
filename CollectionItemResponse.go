package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

type CollectionItemResponse struct {
	Id     string           `json:"id"`
	Name   string           `json:"name"`
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Header []ResponseHeader `json:"header"`
	Body   string           `json:"body"`
}

func (r CollectionItemResponse) BodyIncludePath(request CollectionItemRequest) string {
	dir, file := filepath.Split(request.ShortUrl())
	responseName := strings.Replace(r.Name, " ", "_", 9999)
	file = fmt.Sprintf("%v-%v-%v", strings.ToLower(request.Method), responseName, file)
	return fmt.Sprintf("responses%v%v.json", dir, file)
}

func (r CollectionItemResponse) BodyIncludeString(request CollectionItemRequest) template.HTML {
	return template.HTML(fmt.Sprintf("<!-- include(%v) -->", r.BodyIncludePath(request)))
}

func (r CollectionItemResponse) ContentType() string {
	contentType := ""

	for _, header := range r.Header {
		if strings.ToLower(header.Key) == "content-type" {
			contentType = header.Value
		}
	}

	return contentType
}

func (r CollectionItemResponse) FormattedBody() string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(r.Body), "", "    ")
	if err != nil {
		return r.Body
	}
	return out.String()
}
