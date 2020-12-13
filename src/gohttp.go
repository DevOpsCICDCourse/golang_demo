package main

import (
	"flag"
	"fmt"
	"net/http"
)

var msg = flag.String("message", "Hello, 这是一个Demo页面！", "这是页面的提示消息")

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, *msg)
}

func main() {
	flag.Parse()
	fmt.Println(*msg)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
