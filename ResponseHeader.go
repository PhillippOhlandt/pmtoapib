package main

type ResponseHeader struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
