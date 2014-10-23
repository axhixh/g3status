package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// return power/battery status using information from
// /sys/class/power_supply/BAT0/

func GetPower() string {
	const discharging = '\uf215'
	const charging = '\uf25b'
	var status int
	if isDischarging() {
		status = discharging
	} else {
		status = charging
	}
	return fmt.Sprintf("{\"name\":\"power\", \"full_text\":\"%c %3d%%\"}", status, getCapacity())
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
