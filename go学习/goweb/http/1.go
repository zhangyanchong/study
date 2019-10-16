package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>hello world</h1>")
}

func Nihao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>nihao</h1>")
}

func testPost(w http.ResponseWriter, r *http.Request) { //post
	fmt.Fprintln(w, r.PostFormValue("username"))
	fmt.Fprintln(w, r.PostFormValue("password"))
}

func testGet(w http.ResponseWriter, r *http.Request) { //get
	queryForm := r.URL.Query()
	fmt.Fprintln(w, queryForm["username"][0])
	fmt.Fprintln(w, queryForm["password"][0])
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/nihao", Nihao)
	http.HandleFunc("/test", testPost)
	http.HandleFunc("/testGet", testGet)
	log.Fatal(http.ListenAndServe(":1111", nil))
}
