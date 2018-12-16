package yarl

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/pkg/errors"
)

// Header add the given key value pair to request header.
func (req *Request) Header(k, v string) *Request {
	req.header.Add(k, v)
	return req
}

// Headers
func (req *Request) Headers(v interface{}) *Request {
	switch headers := v.(type) {
	case http.Header:
		req.header = headers
		return req
	}

	rv, rt := reflect.ValueOf(v), reflect.TypeOf(v)
	for rt.Kind() == reflect.Ptr {
		rv, rt = rv.Elem(), rt.Elem()
	}

	switch rt.Kind() {
	case reflect.Struct:
		for i := 0; i < rt.NumField(); i++ {
			if k, ok := rt.Field(i).Tag.Lookup("header"); ok {
				req.header.Add(
					k,
					fmt.Sprintf("%v", rv.Field(i).Interface()),
				)
			}

		}
	case reflect.Map:
		for _, k := range rv.MapKeys() {
			req.header.Add(
				k.String(),
				fmt.Sprintf("%v", rv.MapIndex(k).Interface()),
			)
		}
	default:
		req.err = errors.Errorf("unsupported headers type")
	}

	return req
}

func (req *Request) ContentType(t string) *Request {
	return req.Header("Content-Type", t)
}

func (req *Request) Cookie(c *http.Cookie) *Request {
	req.cookies = append(req.cookies, c)
	return req
}
