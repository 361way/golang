package nux

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/toolkits/file"
)

type Mem struct {
	Buffers   uint64
	Cached    uint64
	MemTotal  uint64
	MemFree   uint64
	SwapTotal uint64
	SwapUsed  uint64
	SwapFree  uint64
	PmemFree  float64
	PmemUsed  float64
	PswapFree float64
	PswapUsed float64 
}

func (this *Mem) String() string {
	return fmt.Sprintf("<MemTotal:%d, MemFree:%d, Buffers:%d, Cached:%d, PmemFree:%.2f, PswapFree:%.2f>", this.MemTotal, this.MemFree, this.Buffers, this.Cached, this.PmemFree, this.PswapFree)
}

var Multi uint64 = 1024

var WANT = map[string]struct{}{
	"Buffers:":   struct{}{},
	"Cached:":    struct{}{},
	"MemTotal:":  struct{}{},
	"MemFree:":   struct{}{},
	"SwapTotal:": struct{}{},
	"SwapFree:":  struct{}{},
}

func MemInfo() (*Mem, error) {
	contents, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	memInfo := &Mem{}

	reader := bufio.NewReader(bytes.NewBuffer(contents))

	for {
		line, err := file.ReadLine(reader)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return nil, err
		}

		fields := strings.Fields(string(line))
		fieldName := fields[0]

		_, ok := WANT[fieldName]
		if ok && len(fields) == 3 {
			val, numerr := strconv.ParseUint(fields[1], 10, 64)
			if numerr != nil {
				continue
			}
			switch fieldName {
			case "Buffers:":
				memInfo.Buffers = val * Multi
			case "Cached:":
				memInfo.Cached = val * Multi
			case "MemTotal:":
				memInfo.MemTotal = val * Multi
			case "MemFree:":
				memInfo.MemFree = val * Multi
			case "SwapTotal:":
				memInfo.SwapTotal = val * Multi
			case "SwapFree:":
				memInfo.SwapFree = val * Multi
			}
		}
	}

	memInfo.SwapUsed = memInfo.SwapTotal - memInfo.SwapFree
	// 增加百分比输出
	memFree := memInfo.MemFree + memInfo.Buffers + memInfo.Cached
	memUsed := memInfo.MemTotal - memFree

	pmemFree := 0.0
	pmemUsed := 0.0
	if memInfo.MemTotal != 0 {
		pmemFree = float64(memFree) * 100.0 / float64(memInfo.MemTotal)
		pmemUsed = float64(memUsed) * 100.0 / float64(memInfo.MemTotal)
	}

	pswapFree := 0.0
	pswapUsed := 0.0
	if memInfo.SwapTotal != 0 {
		pswapFree = float64(memInfo.SwapFree) * 100.0 / float64(memInfo.SwapTotal)
		pswapUsed = float64(memInfo.SwapUsed) * 100.0 / float64(memInfo.SwapTotal)
	}
  memInfo.PmemFree = pmemFree
	memInfo.PmemUsed = pmemUsed
	memInfo.PswapFree = pswapFree
	memInfo.PswapUsed = pswapUsed
 

	return memInfo, nil
}

