Go and Distributed Systems
================================================================================

HTTP
--------------------------------------------------------------------------------

```
$ curl localhost:8000
curl: (7) Failed to connect to localhost port 8000: Connection refused
```

With a web browser

![](images/cant-be-reached.png)


`app.go`

```go
package main

import (
    "net/http"
)

func main() {
    http.ListenAndServe(":8000", nil)
}
```

```bash
$ curl localhost:8000
404 page not found
```

With a web browser

![](images/404.png)


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
⏳
```

```bash
$ curl localhost:8000
```

```bash
$ go run app.go
2022-12-22 12:26:33.497397732 +0100 CET m=+75.067678037
⏳
```

```bash
$ curl localhost:8000
```

```bash
$ go run app.go 
2022-12-22 12:26:33.497397732 +0100 CET m=+75.067678037
2022-12-22 12:28:16.657047211 +0100 CET m=+178.227327515
⏳
```

Via a web browser

![](images/ping.png)

```
$ go run app.go 
2022-12-22 12:29:17.925020503 +0100 CET m=+239.495300807
2022-12-22 12:29:18.128862112 +0100 CET m=+239.699142416
2022-12-22 12:29:18.980003382 +0100 CET m=+240.550283695
⏳
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
$ curl localhost:8001
```


```bash
$ PORT=8001 go run app.go
2022-12-22 12:43:15.285093389 +0100 CET m=+4.132247194
⏳
```

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello world!\n"))
}
```

```bash
$ curl localhost:8000
Hello world!
```

With a web browser

![](images/hello-world-plain-text.png)


```go
func Handler(w http.ResponseWriter, r *http.Request) {
    html := "<p>Hello <b>world!</b></p>\n"
    w.Write([]byte(html))
}
```

```bash
$ curl localhost:8000
<p>Hello <b>world!</b></p>
```

From browser

![](images/hello-world-html.png)

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    IPAddr := r.RemoteAddr
    w.Write([]byte("Hello " + IPAddr + "!\n"))
}
```

```bash
$ curl localhost:8000
Hello 127.0.0.1:48358!
```

```go
func GetIPAddr(r *http.Request) string {
    IPAddrPort := r.RemoteAddr
    if IPAddrPort == "" {
        panic("IP address is undefined")
    }
    IPAddr := strings.Split(IPAddrPort, ":")[0] // 🪲: very brittle
    return IPAddr
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello " + GetIPAddr(r) + "!\n"))
}
```

```bash
$ curl localhost:8000
Hello 127.0.0.1!
```

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "plain/html")
    w.Write([]byte("Hello " + GetIPAddr(r) + "!\n"))
}

func API(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    json := fmt.Sprintf(`{"ip_address": "%s"}`, GetIPAddr(r)) + "\n"
    w.Write([]byte(json))
}

func main() {
    http.HandleFunc("/", Handler)
    http.HandleFunc("/api", API)
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

```bash
$ curl localhost:8000
Hello 127.0.0.1!
$ curl localhost:8000\
Hello 127.0.0.1!
$ curl --head localhost:8000
HTTP/1.1 200 OK
Date: Thu, 22 Dec 2022 15:02:23 GMT
Content-Length: 17
Content-Type: text/plain; charset=utf-8

```

```bash
$ curl localhost:8000/api
{"ip_address": "127.0.0.1"}
$ curl --head localhost:8000/api
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 22 Dec 2022 15:03:37 GMT
Content-Length: 28

```

### Concurrency

The HTTP server handles the requests concurrently

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "plain/html")
    w.Write([]byte("Hello " + GetIPAddr(r) + "!\n"))
    time.Sleep(10 * time.Second)
}
```

```bash
$ curl localhost:8000  # ⏳ Wait ~10 sec
Hello 127.0.0.1!
```

```bash
$ curl localhost:8000 & # Run in the background
$ curl localhost:8000 &
$ curl localhost:8000 &
$ # ⏳ Wait ~10 sec
Hello 127.0.0.1!
Hello 127.0.0.1!
Hello 127.0.0.1!

### 🚧 TODO

  - More client options: Python (requests) and Go

  - More server options: FastAPI & Flask

  - Routing & Query Parameters

  - mDNS ?

  - "Prod:" Compile, 80 as default port, deploy, etc.

  - Concurrency: implicit ("hammer"/"DOS" the server) and implicit.

  - Other protocols & associated go libs reference.
    "Always bet on the web" mostly

    - mDNS

    - mqtt

    - sockets / websockets

    - webRTC


