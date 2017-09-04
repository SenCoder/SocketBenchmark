package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Sample struct {
	Time    int64
	Latency int64
}

func WriteData(s Sample) {
	data, err := json.Marshal(s)
	err = ioutil.WriteFile("sample.json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}
