package main

import (
	"encoding/json"
)

type CollectionItemRequestUrl struct {
	Raw string `json:"raw"`
}

func (u *CollectionItemRequestUrl) UnmarshalJSON(data []byte) error {
	var raw interface{}
	json.Unmarshal(data, &raw)
	switch raw := raw.(type) {
	case string:
		*u = CollectionItemRequestUrl{raw}
	case map[string]interface{}:
		*u = CollectionItemRequestUrl{raw["raw"].(string)}
	}
	return nil
	return nil
}
