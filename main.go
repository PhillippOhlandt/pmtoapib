package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

func getResponseFiles(c Collection) []map[string]string {
	var files []map[string]string

	for _, item := range c.Items {
		if !item.Request.IsExcluded() {
			for _, response := range item.ResponseList() {
				m := map[string]string{}
				m["path"] = response.BodyIncludePath(item.Request)
				m["body"] = response.FormattedBody()
				files = append(files, m)
			}
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

		for _, file := range getResponseFiles(c) {
			writeToFile(
				fmt.Sprintf("%v/%v", filepath.Clean(config.DestinationPath), file["path"]),
				file["body"],
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
