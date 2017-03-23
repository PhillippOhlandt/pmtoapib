package main

import (
	"net/url"
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

func (r CollectionItemRequest) IsExcluded() bool {
	return strings.Contains(r.Description, "pmtoapib_exclude")
}
