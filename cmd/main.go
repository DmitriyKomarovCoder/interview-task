package main

import (
	"bufio"
	"fmt"
	"gorotineCounter/internal/worker"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	inputCh := make(chan string)
	outputCh := make(chan string)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputCh <- scanner.Text()
		}
		close(inputCh)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker.Run(inputCh, outputCh, &wg)
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	for res := range outputCh {
		fmt.Print(res)
	}
	fmt.Printf("Total %d\n", worker.Total)
}
