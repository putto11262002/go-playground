package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	res, err := http.Get("http://localhost:8080/large-file.txt")
	if err != nil {
		fmt.Printf("failed to get file: %v", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	buffer := make([]byte, 1024)

	file, err := os.Create("local/large-file.txt")
	if err != nil {
		fmt.Printf("failed to write to file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	var totalWritten int 

	for {
		nRead, err := res.Body.Read(buffer)

		if err != nil && err != io.EOF {
			
			fmt.Printf("failed to read from response body: %v", err)
			os.Exit(1)
		}

		if nRead == 0 {
			break
		}
		

		nWritten, err := file.Write(buffer[:nRead]);
		
		if err != nil {
			
				fmt.Printf("failed to write to file: %v", err)
				os.Exit(1)
		
		}

		totalWritten += nWritten
	}

	fmt.Printf("wrote %d bytes to %s", totalWritten, "local/large-file.txt")

	


	

}