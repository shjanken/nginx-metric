package metric

import (
	"fmt"
	"os"
	"testing"

	test "com.github.shjanken/nginx-metric/testdata"
)

func TestReadData(t *testing.T) {
	if os.Getenv("BIG_FILE") == "" {
		t.Skip("skip read testdata/access.log file test")
	}
	logfile, err := os.Open("testdata/access.log")
	if err != nil {
		t.Fatalf("open file failure. %v", err)
	}
	fprovider := fileProvider{
		logFile: logfile,
		config:  test.TestConfig,
	}
	ch := make(chan Item, 10)

	go fprovider.ReadData(ch)

	var cnt = 0
	for d := range ch {
		fmt.Println(d.Log)
		cnt++
		fmt.Println("----------------------\n", cnt)
	}
}
