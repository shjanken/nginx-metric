package metric

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/satyrius/gonx"
)

const fileError = "log file is nil or config string is empty"

// FileProvider respect a log file data
type fileProvider struct {
	logFile *os.File
	config  string
}

// NewFileProvider 返回一个从 file 中一行一行读取数据的 DataProvider
func NewFileProvider(file *os.File, config string) DataProvider {
	return &fileProvider{
		logFile: file,
		config:  config,
	}
}

// ReadData read the nginx log data and parse it
func (fprovider *fileProvider) ReadData(ch chan *Item) error {
	if fprovider.logFile == nil || fprovider.config == "" {
		panic(fileError)
	}
	config := strings.NewReader(fprovider.config)
	reader, err := gonx.NewNginxReader(fprovider.logFile, config, "main")
	if err != nil {
		return fmt.Errorf("create file reader failure %w", err)
	}

	for {
		rec, err := reader.Read()
		if err == io.EOF {
			close(ch)
			break
		} else if err != nil {
			ch <- &Item{
				Error: fmt.Errorf("read log failure"),
			}
		} else {
			request := readDataFromGnoxEntry(rec, "request")
			remoteAddr := readDataFromGnoxEntry(rec, "remote_addr")
			remoteUser := readDataFromGnoxEntry(rec, "remote_user")
			timeStr := readDataFromGnoxEntry(rec, "time_local")
			timeLocal, err := time.Parse("02/Jan/2006:15:04:05 -0700", timeStr)
			if err != nil {
				timeLocal = time.Now()
			}
			status, err := strconv.Atoi(readDataFromGnoxEntry(rec, "status"))
			if err != nil {
				status = 0
			}
			bodyBytes, err := strconv.Atoi(readDataFromGnoxEntry(rec, "body_bytes_sent"))
			if err != nil {
				bodyBytes = 0
			}
			httpReferer := readDataFromGnoxEntry(rec, "http_referer")
			httpUserAgent := readDataFromGnoxEntry(rec, "http_user_agent")

			ch <- &Item{
				Log{
					Request:       request,
					RemoteAddr:    remoteAddr,
					RemoteUser:    remoteUser,
					TimeLocal:     timeLocal,
					Status:        statusCode(status),
					BodyBytes:     uint(bodyBytes),
					HTTPReferer:   httpReferer,
					HTTPUserAgent: httpUserAgent,
				},
				nil,
			}
		}
	}

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
