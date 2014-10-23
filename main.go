package main

import (
    "fmt"
    "net"
    "strings"
    "time"
)

func printHeader() {
    fmt.Println(`{"version":1}[`)
}

func printRow() {
    fmt.Printf("[%s,%s,%s,%s,%s],\n", 
        GetPower(),
        getIpAddr(),
        getLocalTime(),
        getTimeAt("America/New_York"),
        getTimeAt("Asia/Kathmandu"))
}

func getBlock(name, instance, text string) string {
    return fmt.Sprintf(`{"name":"%s", "instance":"%s", "full_text":"%s"}`, 
        name, instance, text)
}

func getLocalTime() string {
    const layout = "Mon Jan 02 15:04"
    formattedTime := time.Now().Format(layout)
    return getBlock("time", "local", formattedTime)
}

func getTimeAt(location string) string {
    const layout = "MST 15:04"
    loc,_ := time.LoadLocation(location)
    formattedTime := time.Now().In(loc).Format(layout)
    return getBlock("time", location, formattedTime)
}

func getIpAddr() string {
    addrs, _ := net.InterfaceAddrs()
    for _, addr := range addrs {
        if strings.HasPrefix(addr.String(), "127") {
            continue
        }
        return getBlock("net",addr.Network(), addr.String())
    }
    return getBlock("net","default", "No network")
}

func main() {
    printHeader()
    for ;; {
        printRow()
        time.Sleep(1 * time.Second)
    }
}

