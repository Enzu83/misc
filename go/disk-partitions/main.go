package main

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

func main() {
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, partition := range partitions {
		fmt.Println("Mount point:", partition.Mountpoint)
	}
}
