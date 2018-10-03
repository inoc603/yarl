package yarl

import (
	"fmt"
	"reflect"
)

// Query sets a query item.
func (req *Request) Query(k, v string) *Request {
	if req.hasError() {
		return req
	}

	q := req.url.Query()
	q.Add(k, v)
	req.url.RawQuery = q.Encode()
	return req
}

// Queries add the given values to the query
func (req *Request) Queries(v interface{}) *Request {
	if req.hasError() {
		return req
	}

	q := req.url.Query()
	rv, rt := reflect.ValueOf(v), reflect.TypeOf(v)
	for rt.Kind() == reflect.Ptr {
		rv, rt = rv.Elem(), rt.Elem()
	}

	// TODO: Need to make sure using Sprintf here is correct for all possible types.
	switch rt.Kind() {
	case reflect.Struct:
		for i := 0; i < rt.NumField(); i++ {
			if k, ok := rt.Field(i).Tag.Lookup("query"); ok {
				q.Add(k, fmt.Sprintf("%v", rv.Field(i).Interface()))
			}

		}
	case reflect.Map:
		// TODO: Accept keys other than string? Give an error for non-string keys?
		for _, k := range rv.MapKeys() {
			q.Add(k.String(), fmt.Sprintf("%v", rv.MapIndex(k).Interface()))
		}
	}
	req.url.RawQuery = q.Encode()

	return req
}
