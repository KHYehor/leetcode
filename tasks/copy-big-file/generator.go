//go:build ignore

// Run with: go run generator.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	outputFile = "input.txt"
	targetSize = int64(10 * 1024 * 1024 * 1024) // 10 GiB
	bufferSize = 4 * 1024 * 1024
)

const loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.\n"

func main() {
	if err := generateFile(outputFile, targetSize); err != nil {
		fmt.Fprintf(os.Stderr, "generate %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Printf("generated %s (%d bytes)\n", outputFile, targetSize)
}

func generateFile(path string, size int64) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := file.Close(); err == nil {
			err = closeErr
		}
	}()

	writer := bufio.NewWriterSize(file, bufferSize)
	remaining := size
	text := []byte(loremIpsum)

	for remaining > 0 {
		chunk := text
		if int64(len(chunk)) > remaining {
			chunk = chunk[:remaining]
		}

		written, writeErr := writer.Write(chunk)
		remaining -= int64(written)
		if writeErr != nil {
			return writeErr
		}
		if written != len(chunk) {
			return io.ErrShortWrite
		}
	}

	return writer.Flush()
}
