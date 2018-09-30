# yarl

Yet Another http Request Library in golang. Because why not ?

## Install

```
go get -u -v github.com/inoc603/yarl
```

## Usage

```go
resp, err := yarl.Get("http://google.com").Do()
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

### Setting body

```go
var body struct {
        K string
}

yarl.Get("http://github.com/inoc603").
        Body(&body)
```

### Retry

TODO

### Redirect

TODO

### Unix Socket

```go
resp, err := Get("http://whatever/v1.24/containers/json").
        UnixSocket("/var/run/docker.sock").
        Do()

if err == nil {
        fmt.Println(resp.BodyString())
} else {
        fmt.Println(err)
}
```
