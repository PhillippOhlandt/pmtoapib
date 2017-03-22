package main

type CollectionItemResponse struct {
	Id     string           `json:"id"`
	Name   string           `json:"name"`
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Header []ResponseHeader `json:"header"`
	Body   string           `json:"body"`
}
