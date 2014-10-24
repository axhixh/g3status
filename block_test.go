package main

import "testing"

func TestBlockToJson(t *testing.T) {
	b := &Block{Name: "block", FullText: "full text"}
	json := b.ToJson()

	if json != "{\"name\":\"block\",\"full_text\":\"full text\"}" {
		t.Error("Json doesn't match expected value.", json)
	}
}
