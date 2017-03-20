package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Config struct {
	CollectionPath         string
	DestinationPath        string
	ApibFileName           string
	ForceApibCreation      bool
	ForceResponsesCreation bool
	DumpRequest            string
}

func (c *Config) Init() {
	flag.StringVar(&c.CollectionPath, "collection", "", "Path to the Postman collection export")
	flag.StringVar(&c.CollectionPath, "c", "", "Path to the Postman collection export")

	flag.StringVar(&c.DestinationPath, "destination", "./", "Destination folder path for the generated files")
	flag.StringVar(&c.DestinationPath, "d", "./", "Destination folder path for the generated files")

	flag.StringVar(&c.ApibFileName, "apibname", "", "Set a custom name for the generated .apib file")

	flag.BoolVar(&c.ForceApibCreation, "force-apib", false, "Override existing .apib files")
	flag.BoolVar(&c.ForceResponsesCreation, "force-responses", false, "Override existing response files")

	flag.StringVar(&c.DumpRequest, "dump-request", "", "Output the markup for a single request. (Takes a request name)")

	flag.Parse()
}

type Collection struct {
	Info  CollectionInfo   `json:"info"`
	Items []CollectionItem `json:"item"`
}

type CollectionInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CollectionItem struct {
	Name    string                `json:"name"`
	Request CollectionItemRequest `json:"request"`
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
+ Response 200 (application/json)

    + Headers

            NAME: VALUE

    + Body

            {{ .Request.ResponseBodyIncludeString }}
`

	t := template.New("Item Template")
	t, _ = t.Parse(tpl)

	var doc bytes.Buffer
	t.Execute(&doc, i)
	s := doc.String()

	return template.HTML(s)
}

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

	for key, _ := range m {
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

type RequestHeader struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Disabled    bool   `json:"disabled"`
}

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

func getApibFileContent(c Collection) string {
	tpl :=
		`# Group {{ .Info.Name }}

{{ .Info.Description }}

{{ range .Items }}{{ if not .Request.IsExcluded }}
{{ .Markup }}



{{ end }}{{ end }}
`

	t := template.New("Template")
	t, _ = t.Parse(tpl)

	var doc bytes.Buffer
	t.Execute(&doc, c)
	s := doc.String()

	return s
}

func getResponseFiles(c Collection) []string {
	var files []string

	for _, item := range c.Items {
		if !item.Request.IsExcluded() {
			files = append(files, item.Request.ResponseBodyIncludePath())
		}
	}

	return files
}

func writeToFile(path string, content string, force bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) || force {
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		err := ioutil.WriteFile(path, []byte(content), 0644)
		if err == nil {
			fmt.Printf("Created %v\n", path)
		}
	}
}

func shouldWriteFiles(c Config) bool {
	return c.DumpRequest == ""
}

func main() {

	config := Config{}
	config.Init()

	if config.CollectionPath == "" {
		fmt.Println("No collection file defined!")
		return
	}

	file, _ := ioutil.ReadFile(config.CollectionPath)
	var c Collection
	json.Unmarshal(file, &c)

	apibFileName := strings.Replace(c.Info.Name, " ", "-", -1)

	if config.ApibFileName != "" {
		apibFileName = config.ApibFileName
	}

	apibFile := getApibFileContent(c)

	if shouldWriteFiles(config) {
		writeToFile(
			fmt.Sprintf("%v/%v.apib", filepath.Clean(config.DestinationPath), apibFileName),
			apibFile,
			config.ForceApibCreation,
		)

		for _, path := range getResponseFiles(c) {
			writeToFile(
				fmt.Sprintf("%v/%v", filepath.Clean(config.DestinationPath), path),
				"{}",
				config.ForceResponsesCreation,
			)
		}
	}

	if config.DumpRequest != "" {
		for _, request := range c.Items {
			if request.Name == config.DumpRequest {
				fmt.Println(request.Markup())
			}
		}
	}
}
