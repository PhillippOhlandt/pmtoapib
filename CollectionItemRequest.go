package main

import (
	"fmt"
	"html/template"
	"net/url"
	"path/filepath"
	"sort"
	"strings"
)

type CollectionItemRequest struct {
	Url         string          `json:"url"`
	Method      string          `json:"method"`
	Header      []RequestHeader `json:"header"`
	Body        RequestBody     `json:"body"`
	Description string          `json:"description"`
}

func (r CollectionItemRequest) ShortUrl() string {
	u, _ := url.Parse(r.Url)
	return u.Path
}

func (r CollectionItemRequest) UrlParameterList() []string {
	u, _ := url.Parse(r.Url)

	parameters := []string{}

	m, _ := url.ParseQuery(u.RawQuery)

	for key := range m {
		parameters = append(parameters, key)
	}

	sort.Strings(parameters)

	return parameters
}

func (r CollectionItemRequest) UrlParameterListString() string {
	return strings.Join(r.UrlParameterList(), ",")
}

func (r CollectionItemRequest) ResponseBodyIncludePath() string {
	dir, file := filepath.Split(r.ShortUrl())
	file = fmt.Sprintf("%v-%v", strings.ToLower(r.Method), file)
	return fmt.Sprintf("responses%v%v.json", dir, file)
}

func (r CollectionItemRequest) ResponseBodyIncludeString() template.HTML {
	return template.HTML(fmt.Sprintf("<!-- include(%v) -->", r.ResponseBodyIncludePath()))
}

func (r CollectionItemRequest) IsExcluded() bool {
	return strings.Contains(r.Description, "pmtoapib_exclude")
}
