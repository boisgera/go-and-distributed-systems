Go and Distributed Systems
================================================================================

Go
--------------------------------------------------------------------------------

**TODO:**

  - go tour reference

  - selected examples (to make the following HTTP Section understandable
    first and foremost)

HTTP
--------------------------------------------------------------------------------

### Hello World

`app.go`
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now())
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8000", nil)
}
```

```bash
$ go run app.go
```

```bash
$ curl localhost:8000
```

```bash
$ go run app.go
2022-12-22 12:26:33.497397732 +0100 CET m=+75.067678037
```

```bash
$ curl 127.0.0.1:8000
```

```
$ go run app.go 
2022-12-22 12:26:33.497397732 +0100 CET m=+75.067678037
2022-12-22 12:28:16.657047211 +0100 CET m=+178.227327515
```

Via a web browser

![](images/hello-world-plain-text.png)

```
$ go run app.go 
2022-12-22 12:29:17.925020503 +0100 CET m=+239.495300807
2022-12-22 12:29:18.128862112 +0100 CET m=+239.699142416
2022-12-22 12:29:18.980003382 +0100 CET m=+240.550283695
```

```go
func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
```

```go
$ go run app.go
``

```go
$ go run app.go 
panic: listen tcp :8000: bind: address already in use
```

```go
func main() {
	http.HandleFunc("/", Handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
```

```
$ PORT=8001 go run app.go
```

```
$ curl 127.0.0.1:8001
```


```bash
$ PORT=8001 go run app.go
2022-12-22 12:43:15.285093389 +0100 CET m=+4.132247194
```

----

```go
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!\n"))
}
```

```bash
$ curl 127.0.0.1:8000
Hello world!
```

```go
func Handler(w http.ResponseWriter, r *http.Request) {
	html := "<p>Hello <b>world!</b></p>\n"
	w.Write([]byte(html))
}
```

```bash
$ curl 127.0.0.1:8000
<p>Hello <b>world!</b></p>
```

From browser

![](images/hello-world-html.png)






  2. Connect to it via curl, browser (and LATER, Python (requests) and Go)

  3. Return "Hello world" HTML instead of plain text

  4. Return "Hello world" + IP adress + date

  5. Do it in JSON

  6. Routing & Query Parameters

     (Mention FastAPI & Flask alternative)

  7. mDNS ?

  8. Compile, 80 as default port, deploy, etc.

Q: How can I introduce concurrency fast here ? Implicitly with a slow request
and show that the HTTP Server IS concurrent. But explicitly? Start a 
long-running process that collects a log of the connections and give the list
back? So that would be a 4.5 ; yes.

Moar
--------------------------------------------------------------------------------

Concepts and Go libs

  - mDNS

  - mqtt

  - sockets / websockets

  - webRTC


