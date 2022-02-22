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
	if path == "/" {
		if _, err := os.Stat("./index.html"); err == nil {
			http.ServeFile(w, r, "index.html")
		} 
		else {
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