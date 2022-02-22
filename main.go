package main

import (
	"github.com/julienschmidt/httprouter"
    "fmt"
    "net/http"
    "log"
	"os"
	"io/ioutil"
)

func path(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")
	path_local := "./" + path

	// check that path does not goes to parent directory
	// breake path and calculate where it goes to
	// if it goes to parent directory, return error
	// if it goes to child directory, return file

	// look for ../
	for i := 0; i < len(path); i++ {
		if path[i] == '.' && path[i+1] == '.' && path[i+2] == '/' {
			w.Write([]byte("404"))
		}
	}

	if path == "/" {
		if _, err := os.Stat("./index.html"); err == nil {
			http.ServeFile(w, r, "index.html")
		} else {
			w.Write([]byte("404"))
		}
	} else {
	if _, err := os.Stat(path_local); err == nil {
		if info, err := os.Stat(path_local); err == nil {
			if info.IsDir() {
				files, err := ioutil.ReadDir("." + path)
				if err != nil {
					log.Fatal(err)
				}
				for _, file := range files {
					if info, err := os.Stat(path_local); err == nil {
						if info.IsDir() {
							fmt.Fprintf(w, "%s\n", file.Name())
						}
					}
				}
			} else {
				http.ServeFile(w, r, "." + path)
			}

		}
	} else {
		fmt.Fprintf(w, "File not found ")
	}
	}
}

func main() {
    router := httprouter.New()
	port := os.Args[1]
	fmt.Println("WEB SERVER IS RUNNING ON PORT ",port)
	router.GET("/*path", path)

    log.Fatal(http.ListenAndServe(":" + port , router))
}