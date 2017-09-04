package util

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func MemInfo() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	log.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	log.Println(v)
}

func CpuInfo() {
	percents, err := cpu.Percent(time.Second*10, false)
	if err != nil {
		log.Println(percents)
	}
}
