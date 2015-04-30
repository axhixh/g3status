package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func GetPower() *Block {
	const discharging = '\uf215'
	const charging = '\uf25b'

	battery := getBattery()

	var status int
	if isDischarging(battery) {
		status = discharging
	} else {
		status = charging
	}
	capacity := getCapacity(battery)
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

func getBattery() string {
	files, err := ioutil.ReadDir("/sys/class/power_supply")
	if err != nil {
		return "/sys/class/power_supply/BAT0"
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "BAT") {
			return "/sys/class/power_supply/" + file.Name()
		}
	}
	return "/sys/class/power_supply/BAT0"
}

func getCapacity(battery string) int {
	value, err := read(battery + "/capacity")
	if err != nil {
		return -1
	}

	capacity, err := strconv.Atoi(value)
	if err != nil {
		return -1
	}
	return capacity
}

func isDischarging(battery string) bool {
	status, err := read(battery + "/status")
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
