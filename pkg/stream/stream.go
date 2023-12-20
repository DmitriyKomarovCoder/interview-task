package stream

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func UrlStream(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func FileStream(fileName string) ([]byte, error) {
	inputStream, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := inputStream.Close(); closeErr != nil {
			log.Fatalf("err: can't close file %s, with error: %v", fileName, fmt.Errorf("%w", err))
		}
	}()

	body, err := io.ReadAll(inputStream)
	if err != nil {
		return nil, err
	}

	return body, nil
}
