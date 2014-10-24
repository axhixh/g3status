package main

import (
	"encoding/json"
)

// Represents a block on the status bar. Does not
// include all the available properties like colour etc
type Block struct {
	Name     string `json:"name"`
	Instance string `json:"instance,omitempty"`
	FullText string `json:"full_text"`
}

func (block *Block) ToJson() string {
	v, err := json.Marshal(block)
	if err != nil {
		return ""
	}
	return string(v)
}
