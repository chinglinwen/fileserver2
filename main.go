package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
		fmt.Println("Fileserver2 version=1.1.0, 2017-1-6")
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
	http.HandleFunc("/", detector)

	err := http.ListenAndServe(":"+port, nil)
	checkError(err)
}

func detector(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.RequestURI, "uploadapi") {
		uploadHandler(w, r)
		return
	}
	// print logs
	ip := strings.Split(r.RemoteAddr, ":")[0]
	log.Println(ip, r.RequestURI, "visited")

	if strings.HasSuffix(r.RequestURI, "upload") {
		uploadPageHandler(w, r)
		return
	}
	http.FileServer(http.Dir(path)).ServeHTTP(w, r)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
