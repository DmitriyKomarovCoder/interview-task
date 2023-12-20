package worker

import (
	"fmt"
	stream "gorotineCounter/pkg/stream"
	"log"
	"strings"
	"sync"
	"sync/atomic"
)

const entries = "Go"

var Total int32

func countGo(content []byte, entries string) int {
	count := 0
	for i := 0; i+len(entries) <= len(content); i++ {
		if string(content[i:i+len(entries)]) == entries {
			count++
		}
	}
	return count
}

func Run(inputCh chan string, outputCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range inputCh {
		var err error
		var body []byte

		if strings.HasPrefix(line, "http://") || strings.HasPrefix(line, "https://") {
			body, err = stream.UrlStream(line)
			if err != nil {
				log.Printf("err: can't get url body %s, with error: %v", line, fmt.Errorf("%w", err))
				continue
			}
		} else if line != "" {
			body, err = stream.FileStream(line)
			if err != nil {
				log.Printf("err: can't get file body %s, with error: %v", line, fmt.Errorf("%w", err))
				continue
			}
		}
		count := countGo(body, entries)
		atomic.AddInt32(&Total, int32(count))

		select {
		case outputCh <- fmt.Sprintf("Count for %s: %d\n", line, count):
		default:
			return
		}
	}
}
