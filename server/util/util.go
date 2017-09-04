package util

import (
    "log"
    "github.com/shirou/gopsutil/mem"
)

func memInfo() {
    v, _ := mem.VirtualMemory()

    // almost every return value is a struct
    log.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

    // convert to JSON. String() is also implemented
    log.Println(v)
}
