package main

import (
	"fmt"
	"io"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	fmt.Println("µµµ", r.RemoteAddr)
	fmt.Println("***", r.Header.Get("X-Forwarded-For"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	fmt.Println("got a request")
	io.WriteString(w, `
	<html>
  	    <body>
	        <p>Hello <b>world!</b></p>
	    </body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8000", nil)
}
