package main

type CollectionItemResponses []CollectionItemResponse

func (r CollectionItemResponses) Len() int {
	return len(r)
}

func (r CollectionItemResponses) Less(i, j int) bool {
	return r[i].Code < r[j].Code
}

func (r CollectionItemResponses) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
