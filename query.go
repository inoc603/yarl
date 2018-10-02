package yarl

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

	return req
}
