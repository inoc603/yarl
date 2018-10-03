# yarl

[![GoDoc](https://godoc.org/github.com/inoc603/yarl?status.svg)](http://godoc.org/github.com/inoc603/yarl)

Yet Another http Request Library in golang. Because why not ?

**Work In Progress**

## Install

```
go get -u -v github.com/inoc603/yarl
```

## Usage

```go
resp := yarl.Get("http://example.com").Do()

if resp.Error() == nil {
        fmt.Println(resp.StatusCode())
        // Get response body as a reader
        io.Copy(os.Stdout, resp.Body())
        // Get response as a string or bytes
        fmt.Println(resp.BodyString(), resp.BodyBytes())
}

// Marshal response body to a struct
var body struct {
        K string `json:"k"`
}

resp, err := yarl.Get("http://example.com").
        DoMarshal(&body)

if err != nil {
        // Response body can still be used if marshalling failed
        fmt.Printf("error: %v; body: %s", err, resp.BodyString())
}
```

### Setting headers

```go
yarl.Get("http://github.com/inoc603").
        Header("k1", "v1").
        Headers(map[string]string{
                "k2": "v2",
        }).
        Headers(struct {
                K3 string `header:"k3"`
        }{"v3"})
```

### Setting query

```go
yarl.Get("http://github.com/inoc603").
        Query("k1", "v1").
        Queries(map[string]interface{}{
                "k2": "v2",
                "k3": 3,
        }).
        Queries(struct {
                K4 string `query:"k4"`
                K5 int    `query:"k5"`
        }{"v4", 5})
```

### JSON body

```go
// From a struct
body := struct {
        K string `json:"k"`
}{"value"}
yarl.Post("http://github.com/inoc603").
        Body(&body)

// From a map
mapBody := map[string]interface{}{
        "key": "value",
}
yarl.Post("http://github.com/inoc603").
        Body(mapBody)

// From a string or bytes
yarl.Post("http://github.com/inoc603").
        Body(`{"key": { "nested": 1 }}`)

yarl.Post("http://github.com/inoc603").
        Body([]byte(`{"key": { "nested": 1 }}`))
```

### Multipart Body

```go
yarl.Post("http://github.com/inoc603").
        Multipart().
        File("./file1.txt").
        File("./file2.txt", "field_name").
        FileFromReader(bytes.NewBuffer(content), "file3.txt", "field_name_2").
        Do()
```

### Retry

```go
yarl.Get("http://example.com").
        Retry(3, time.Second)
```

### Redirect

TODO

### Unix Socket

```go
yarl.Get("http://whatever/v1.24/containers/json").
        UnixSocket("/var/run/docker.sock").
        Do()
```

### Reuse configurations

TODO: Make reusing reqeust thread-safe

```go
v1 := yarl.New("http://example.com").
        BasePath("/api/v1")

// following calls will reuse v1
v1.Get("/example").Do()
v1.Post("/user/%d", 1).Body(body).Do()

// TODO
```
