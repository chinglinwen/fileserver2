package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
)

var (
	path string
	port string
)

func main() {
	flag.StringVar(&port, "port", "9000", "Port number")
	flag.StringVar(&path, "path", ".", "File server path")
	version := flag.Bool("version", false, "Show version")
	author := flag.Bool("author", false, "Show author")

	verbose := flag.Bool("v", false, "output to logfile ( default stdout )")
	logFile := flag.String("logfile", "fs.log", "log filename and path")
	logMaxSize := flag.Int("logmaxsize", 500, "log max size(megabytes)")
	logMaxAge := flag.Int("logmaxage", 28, "log max age (days)")
	logMaxBackups := flag.Int("logmaxbackups", 3, "log max backups number")

	flag.Parse()

	//Display version info.
	if *version {
		fmt.Println("Fileserver2 version=1.3.0, 2020-7-7")
		os.Exit(0)
	}

	//Display author info.
	if *author {
		fmt.Println("Author is Wen Zhenglin")
		os.Exit(0)
	}

	if *verbose {
		log.SetOutput(&lumberjack.Logger{
			Filename:   *logFile,
			MaxSize:    *logMaxSize, // megabytes
			MaxBackups: *logMaxBackups,
			MaxAge:     *logMaxAge, //days
		})
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
