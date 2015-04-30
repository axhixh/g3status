package main

import (
	"encoding/json"
)

// Represents a block on the status bar.
type Block struct {
	Name     string `json:"name"`
	Instance string `json:"instance,omitempty"`
	FullText string `json:"full_text"`
	Color    string `json:"color,omitempty"`
}

func (block *Block) ToJson() string {
	v, err := json.Marshal(block)
	if err != nil {
		return ""
	}
	return string(v)
}
