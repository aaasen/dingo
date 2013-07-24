package dingo

import (
	"fmt"
	"net/http"
)

func Hello() {
	fmt.Println("hey!")
}

func loadFile(c chan string, path string) {
	// content, _ := ioutil.ReadFile(path)
	// c <- string(content)
	c <- "concurrent"
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	c := make(chan string)
	go loadFile(c, r.URL.Path[1:])
	fmt.Fprintf(w, string(<-c))
}

func Run() {
	http.HandleFunc("/", staticHandler)
	http.ListenAndServe(":"+Port, nil)
}
