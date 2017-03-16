package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

func TestStandardOutputIsCorrect(t *testing.T) {
	file, _ := ioutil.ReadFile("./example/pmtoapib example.postman_collection.json")
	var c Collection
	json.Unmarshal(file, &c)

	apibFile := getApibFileContent(c)
	expectedApibFile, _ := ioutil.ReadFile("./example/docs/pmtoapib-example.apib")

	if apibFile != string(expectedApibFile) {
		t.Error("Generated apib file didn't match with the expected output")
		apibFileSlice := strings.Split(apibFile, "\n")
		expectedApibFileSlice := strings.Split(string(expectedApibFile), "\n")

		for index, text := range apibFileSlice {
			if text != expectedApibFileSlice[index] {
				t.Logf("Expected '%v' to be '%v' in line %v", text, expectedApibFileSlice[index], index)
			}
		}


	}
}