package main

import (
	"bytes"
	"html/template"
	"sort"
)

type CollectionItem struct {
	Name      string                  `json:"name"`
	Request   CollectionItemRequest   `json:"request"`
	Responses CollectionItemResponses `json:"response"`
}

func (i CollectionItem) Markup() template.HTML {
	tpl :=
		`## {{ .Name }} [{{ .Request.ShortUrl }}{{ if .Request.UrlParameterListString }}{?{{ .Request.UrlParameterListString }}}{{ end }}]

### {{ .Name }} [{{ .Request.Method }}]

{{ if .Request.Description }}{{ .Request.Description }}{{ else }}DESCRIPTION{{ end }}
{{ if .Request.UrlParameterList }}
+ Parameters

    {{ range .Request.UrlParameterList }}+ {{ . }} (string, required) - DESCRIPTION
    {{ end }}{{ end }}
+ Request

    + Headers
            {{ range .Request.Header }}{{ if not .Disabled }}
            {{ .Key }}: {{ .Value }}{{ end }}{{ end }}
    {{ if .Request.Body.Raw }}
    + Body

    	    {{ .Request.Body.RawString }}
    {{ end }}
{{ .ResponseSectionMarkup }}
`

	t := template.New("Item Template")
	t, _ = t.Parse(tpl)

	var doc bytes.Buffer
	t.Execute(&doc, i)
	s := doc.String()

	return template.HTML(s)
}

func (i CollectionItem) ResponseSectionMarkup() template.HTML {
	tpl :=
		`{{ range .ResponseList }}
+ Response {{ .Code }}{{ if .ContentType }} ({{ .ContentType }}){{ end }}

    + Headers
            {{ range .Header }}{{ if not .Hidden }}
            {{ .Key }}: {{ .Value }}{{ end }}{{ end }}

    + Body

            {{ .BodyIncludeString $.Request }}
{{ end }}`

	t := template.New("Response Section Template")
	t, _ = t.Parse(tpl)

	var doc bytes.Buffer
	t.Execute(&doc, i)
	s := doc.String()

	return template.HTML(s)
}

func (i CollectionItem) ResponseList() CollectionItemResponses {
	responses := CollectionItemResponses{}

	dummyTwoHundredResponse := CollectionItemResponse{
		Id:     "1",
		Name:   "200",
		Status: "OK",
		Code:   200,
		Header: []ResponseHeader{
			{
				Key:         "Content-Type",
				Value:       "application/json",
				Name:        "Content-Typ",
				Description: "The mime type of this content",
			},
			{
				Key:         "NAME",
				Value:       "VALUE",
				Name:        "NAME",
				Description: "Dummy Header",
			},
		},
		Body: "{}",
	}

	if len(i.Responses) == 0 {
		responses = append(responses, dummyTwoHundredResponse)
		return responses
	}

	responses = i.Responses

	hasTwoHundred := false

	for _, response := range responses {
		if response.Code == 200 {
			hasTwoHundred = true
		}
	}

	if !hasTwoHundred {
		responses = append(responses, dummyTwoHundredResponse)
	}

	sort.Sort(responses)

	return responses
}
