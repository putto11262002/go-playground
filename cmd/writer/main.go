package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	// "time"
)


const (
	defaultSize = 1024
	defaultOut = "output.txt"
)

func createFile(path string) (*os.File, error) {
	return os.Create(path)
}

func writeRandAscii(writer io.WriteCloser, size int, progressCh chan<- float32) (int, error) {
	defer func(){
		writer.Close()
		close(progressCh)
	}()

	var written int

	bufferSize := 1024

	buffer := make([]byte, bufferSize)

	for written < size  {
		var toWrite int
		if size - bufferSize < written {
			toWrite = size - written
		} else {
			toWrite = bufferSize	
		}

		for i := 0; i < int(toWrite); i++ {
			buffer[i] = byte(rand.Intn(95) + 32)
		}

		n, err := writer.Write(buffer[:toWrite])
		if err != nil {
			return written, err
		}
		written += n
		
		if progressCh != nil {
			progressCh <- float32(written) / float32(size)
		}
		
	}
	return written, nil
}



func main() {
	size := flag.Int("size", 1024, "file size in bytes")

	outputPath := flag.String("output", defaultOut, "output file path")

	flag.Parse()

	file, err := createFile(*outputPath)

	if err != nil {
		fmt.Printf("failed to create file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	progressCh := make(chan float32)

	go func(){
		for progress := range progressCh {
			fmt.Printf("\r%.2f%%", progress * 100)
		}
	}()

	
	written, err := writeRandAscii(file, *size, progressCh)


	if err != nil && written == 0 {
		fmt.Printf("\nfailed to writg to file: %v", err)
	}else if err != nil && written > 0 {
		fmt.Printf("\nfailed to complete writing to file: %v", err)
	}

	fmt.Printf("\nSuccessfully wrote %d bytes to %s\n", written, *outputPath)
}
