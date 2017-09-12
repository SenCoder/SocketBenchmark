package util

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Collector struct {
	f *os.File
}

type DataSample struct {
	Time    int64
	Latency int64
	Count   int64
}

func (c *Collector) OpenFile(filename string, perm os.FileMode) {

	var err error
	if c.f == nil {
		c.f, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
		if err != nil {
			log.Println(err)
		}
	}
}

func (c *Collector) CloseFile() {
	if c.f != nil {
		c.f.Close()
		c.f = nil
	}
}

func (c *Collector) Sample(d DataSample) {

	dataStr := strconv.FormatInt(d.Time, 10) + " " + strconv.FormatInt(d.Latency, 10) + " " + strconv.FormatInt(d.Count, 10) + "\n"

	c.f.Write([]byte(dataStr))

}

func Rs(length int) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	buffer := bytes.NewBufferString("")
	for i := 0; i < length; i++ {
		isLetter := r.Intn(2)
		if isLetter > 0 {
			letter := r.Intn(52)
			if letter < 26 {
				letter += 97
			} else {
				letter += 65 - 26
			}
			buffer.WriteString(string(letter))
			//buffer.WriteString(fmt.Sprintf("%c", letter))
		} else {
			buffer.WriteString(strconv.Itoa(r.Intn(10)))
		}
	}
	return buffer.Bytes()
}
