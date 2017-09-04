package util

import (
	"io"
	"log"
	"os"
	"strconv"
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

func (sample DataSample) toString() string {
	return strconv.FormatInt(sample.Time, 10) + " " + strconv.FormatFloat(sample.CpuPerc, 'g', 5, 32) + " " +
		strconv.FormatUint(sample.MemUsed, 10) + " " + strconv.FormatFloat(sample.MemPerc, 'g', 5, 32) + "\n"
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

func Sample() {

	sample := DataSample{Time: time.Now().Unix(), CpuPerc: CpuInfo()}
	sample.MemUsed, sample.MemPerc = MemInfo()

	err := WriteFile("sample.json", []byte(sample.toString()), 0644)

	if err != nil {
		log.Println(err)
	}
}

func MemInfo() (uint64, float64) {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	log.Printf("Used:%v, UsedPercent:%f%%\n", v.Used/1024/1024, v.UsedPercent)

	// convert to JSON. String() is also implemented
	log.Println(v)

	return v.Used / 1024 / 1024, v.UsedPercent
}

func CpuInfo() float64 {
	percents, err := cpu.Percent(time.Second*5, false)
	if err != nil {
		log.Println(percents)
	}
	return percents[0]
}

func CollectData() {
	go func() {
		for {
			Sample()
		}
	}()
}
