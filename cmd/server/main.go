package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main(){
	flag.Parse()

	port := flag.String("port", "8080", "port to serve on")
	filePath := flag.Arg(0)
	

	fmt.Println(filePath)
	if filePath == "" {
		fmt.Println("missing file path")
		os.Exit(1)
	}

	http.Handle("/", http.FileServer(http.Dir(filePath)))

	http.ListenAndServe(":8080", http.FileServer(http.Dir(filePath)))

	fmt.Printf("serving %s at ::%s", filePath, *port)
}