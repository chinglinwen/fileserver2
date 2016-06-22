package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	path string
	port string
)

func main() {
	flag.StringVar(&port, "port", "8000", "Port number.")
	flag.StringVar(&path, "path", ".", "File server path.")
	version := flag.Bool("v", false, "Show version.")
	author := flag.Bool("author", false, "Show author.")

	flag.Parse()

	//Display version info.
	if *version {
		fmt.Println("Fileserver2 version=1.0, 2016-6-22")
		os.Exit(0)
	}

	//Display author info.
	if *author {
		fmt.Println("Author is Wen Zhenglin")
		os.Exit(0)
	}

	if len(os.Args) == 2 {
		argsPath := os.Args[1]
		if argsPath != "" {
			fmt.Println(argsPath)
			path = argsPath
		}
	}

	http.HandleFunc("/upload", uploadHandler)

	http.Handle("/", http.FileServer(http.Dir(path)))

	err := http.ListenAndServe(":"+port, nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
