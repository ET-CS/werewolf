package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

// Get app folder
func getAppDir() string {
	// Get current folder or die
	dir, patherr := filepath.Abs(filepath.Dir(os.Args[0]))
	if patherr != nil {
		log.Fatal(patherr)
	}
	return dir
}

var (
    // Directory where 'werewolf' is currently running in
    projectDir = getAppDir()
)

// Map of all htmls found on walk
var m map[string]string

func visit(path string, f os.FileInfo, err error) error {
	// relative path without projectDir
	rpath := path[len(projectDir):]
	// Check if file is html file
	isHTML := strings.HasSuffix(path, ".html")
	if isHTML {
		// get only file from path
		s := strings.Split(rpath, "/")
		// get filename from path (index.min.html or index.html)
		fn := s[len(s)-1]
		// get filename without extension (index.min or index)
		fn = fn[:len(fn)-5]
		// remove .min if minified
		isMinifiedHTML := strings.HasSuffix(fn, ".min")
		if isMinifiedHTML {
			// remove .min from filename (index)
			fn = fn[:len(fn)-4]
		}
		// TODO get the relative path `/post/` from `/post/1.html` so it will be `/post/1` in the end and not `/1`
		// TODO replace current map to hold struct with {path, rpath, fn}
		m[fn] = path
	}
	return nil
}

// Here everything starts
func main() {
	m = make(map[string]string)
	// Find html files to serve
	err := filepath.Walk(projectDir, visit)
	if err != nil {
		log.Fatal("Walk: ", err)
	}

	// Create muxxer
	r := mux.NewRouter()
	r.StrictSlash(true)
	http.Handle("/", r)

	for key, value := range m {
		// index is /
		if key == "index" {
			key = ""
		}
		// load file into html string
		fc, err := ioutil.ReadFile(value)
		if err != nil {
			log.Fatal("ioutil: ", err)
		}
		html := string(fc)
		// get relative path with projectDir
		rpath := value[len(projectDir):]
		fmt.Printf("serving " + rpath + " on /" + key + "\n")
		// create handler to serve html
		r.HandleFunc("/"+key, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, html)
		}).Methods("GET")

	}

	// Serve static
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	// Starting server
	port := "8585"
	fmt.Println("Starting werewolf server on port: " + port + "...")
	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
