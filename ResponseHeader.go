package main

import "strings"

type ResponseHeader struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h ResponseHeader) Hidden() bool {
	return strings.ToLower(h.Key) == "content-type"
}
