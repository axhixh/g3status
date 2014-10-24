package main

import (
	"fmt"
	"time"
)

func printHeader() {
	fmt.Println(`{"version":1}[`)
}

func printRow() {
	fmt.Printf("[%s,%s,%s,%s,%s,%s],\n",
		GetDisk().ToJson(),
		GetPower().ToJson(),
		GetIpAddr().ToJson(),
		GetLocalTime().ToJson(),
		GetTimeAt("America/New_York").ToJson(),
		GetTimeAt("Asia/Kathmandu").ToJson())
}

func main() {
	printHeader()
	for {
		printRow()
		time.Sleep(1 * time.Second)
	}
}
