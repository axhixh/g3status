package main

import (
	"fmt"
	"syscall"
)

func GetDisk() *Block {
	return &Block{"disk", "/", humanize(diskUsage("/"))}
}

func diskUsage(path string) float64 {
	buf := new(syscall.Statfs_t)
	syscall.Statfs("/", buf)
	return float64(buf.Bsize) * float64(buf.Bfree)
}

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
)

var units = []uint64{TB, GB, MB, KB}
var labels = map[uint64]string{
	KB: "KiB",
	MB: "MiB",
	GB: "GiB",
	TB: "TiB",
}

func humanize(value float64) string {
	for _, unit := range units {
		size := float64(unit)
		if value > size {
			return fmt.Sprintf("%.1f %s", value/size, labels[unit])
		}
	}
	return fmt.Sprintf("%f B", value)
}
