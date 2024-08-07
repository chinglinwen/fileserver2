package main

import (
	"encoding/base64"
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

	auth := flag.Bool("auth", false, "enable basic auth")
	downloadAuth := flag.Bool("downloadauth", false, "enable download basic auth")
	user := flag.String("user", "admin", "user name for basic auth")
	password := flag.String("password", "Admin@Qax123_@", "password for basic auth")

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

	s := server{auth: *auth, downloadAuth: *downloadAuth, user: *user, pass: *password}
	err := http.ListenAndServe(":"+port, s)
	checkError(err)
}

type server struct {
	auth         bool
	downloadAuth bool
	user         string
	pass         string
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// print logs
	ip := strings.Split(r.RemoteAddr, ":")[0]

	if strings.HasSuffix(r.RequestURI, "uploadapi") {
		if s.auth {
			log.Printf("uploadapi(auth), ip: %v, uri: %v, visited", ip, r.RequestURI)
			BasicAuthMiddleware(http.Handler(http.HandlerFunc(uploadHandler)), s.user, s.pass).ServeHTTP(w, r)
			return
		}
		log.Printf("uploadapi(noauth), ip: %v, uri: %v, visited", ip, r.RequestURI)
		uploadHandler(w, r)
		return
	}
	if strings.HasSuffix(r.RequestURI, "upload") {
		if s.auth {
			log.Printf("uploadpage(auth), ip: %v, uri: %v, visited", ip, r.RequestURI)
			BasicAuthMiddleware(http.Handler(http.HandlerFunc(uploadPageHandler)), s.user, s.pass).ServeHTTP(w, r)
			return
		}
		log.Printf("uploadpage(noauth), ip: %v, uri: %v, visited", ip, r.RequestURI)
		uploadPageHandler(w, r)
		return
	}
	if s.downloadAuth {
		log.Printf("dwonload(auth), ip: %v, uri: %v, visited", ip, r.RequestURI)
		BasicAuthMiddleware(http.FileServer(http.Dir(path)), s.user, s.pass).ServeHTTP(w, r)
		return
	}
	log.Printf("dwonload(noauth), ip: %v, uri: %v, visited", ip, r.RequestURI)
	http.FileServer(http.Dir(path)).ServeHTTP(w, r)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// BasicAuthMiddleware is a middleware that provides basic auth
func BasicAuthMiddleware(next http.Handler, username, password string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Printf("auth failed, no header\n")
			return
		}

		// Check if the Authorization header is in the correct format
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Printf("auth failed, invalid header\n")
			return
		}

		// Decode the base64 encoded credentials
		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Printf("auth failed, not base64\n")
			return
		}

		// Check the username and password
		creds := strings.SplitN(string(decoded), ":", 2)
		if len(creds) != 2 || creds[0] != username || creds[1] != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Printf("auth failed, incorrect password\n")
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}
