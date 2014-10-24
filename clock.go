package main

import "time"

func GetLocalTime() *Block {
	const layout = "Mon Jan 02 15:04"
	formattedTime := time.Now().Format(layout)
	return &Block{"time", "local", formattedTime}
}

func GetTimeAt(location string) *Block {
	const layout = "MST 15:04"
	loc, _ := time.LoadLocation(location)
	formattedTime := time.Now().In(loc).Format(layout)
	return &Block{"time", location, formattedTime}
}
