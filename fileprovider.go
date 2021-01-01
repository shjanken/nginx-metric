package metric

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/satyrius/gonx"
)

const fileError = "log file is nil or config string is empty"

// FileProvider respect a log file data
type fileProvider struct {
	logFile *os.File
	config  string
	ch      chan Item
}

// ReadData read the nginx log data and parse it
func (fprovider *fileProvider) ReadData(ch chan Item) {
	if fprovider.logFile == nil || fprovider.config == "" {
		panic(fileError)
	}
	config := strings.NewReader(fprovider.config)
	reader, err := gonx.NewNginxReader(fprovider.logFile, config, "main")
	if err != nil {
		// 如果创建 NginxReader 失败， 直接退出程序
		log.Fatalf("create nginx reader failure, %v", err)
	}

	for {
		rec, err := reader.Read()
		if err == io.EOF {
			close(ch)
			break
		} else if err != nil {
			//TODO 处理错误
		}

		request := readDataFromGnoxEntry(rec, "request")

		ch <- Item{
			Log{Request: request},
			nil,
		}
	}
}

func (fprovider fileProvider) Save() error {
	return nil
}

// read the data from gnox.Entry
// if gnox.Entry.Field() return error, return ""
func readDataFromGnoxEntry(entry *gonx.Entry, field string) string {
	val, err := entry.Field(field)
	if err != nil {
		return ""
	}
	return val
}
