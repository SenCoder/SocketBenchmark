package util

import (
	"encoding/json"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type DataSample struct {
	Time    int64
	CpuPerc float64
	MemUsed uint64
	MemPerc float64
}

func Sample() {

	sample := DataSample{Time: time.Now().Unix(), CpuPerc: CpuInfo()}
	sample.MemUsed, sample.MemPerc = MemInfo()
	data, _ := json.Marshal(sample)

	err = ioutil.WriteFile("sample.json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func MemInfo() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	log.Printf("Used:%v, UsedPercent:%f%%\n", v.Used/1024/1024, v.UsedPercent)

	// convert to JSON. String() is also implemented
	log.Println(v)

	return v.Used / 1024 / 1024, v.UsedPercent
}

func CpuInfo() {
	percents, err := cpu.Percent(time.Millisecond*10, false)
	if err != nil {
		log.Println(percents)
	}
	return percents[0]
}
