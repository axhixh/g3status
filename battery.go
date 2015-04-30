package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// return power/battery status using information from
// /sys/class/power_supply/BAT0/

func GetPower() *Block {
	const discharging = '\uf215'
	const charging = '\uf25b'
	var status int
	if isDischarging() {
		status = discharging
	} else {
		status = charging
	}
	capacity := getCapacity()
	var color string
	if capacity > 20 {
		color = "#00ff00"
	} else {
		color = "#ff0000"
	}

	return &Block{
		Name:     "power",
		FullText: fmt.Sprintf("%c %3d%%", status, capacity),
		Color:    color}
}

func getCapacity() int {
	value, err := read("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		return -1
	}

	capacity, err := strconv.Atoi(value)
	if err != nil {
		return -1
	}
	return capacity
}

func isDischarging() bool {
	status, err := read("/sys/class/power_supply/BAT0/status")
	if err != nil {
		return false
	}
	return "Discharging" == status
}

func read(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}
