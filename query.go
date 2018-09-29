package yarl

func (req *Request) Query(k, v string) *Request {
	if req.req == nil || req.err != nil {
		return req
	}

	req.req.URL.Query().Add(k, v)
	return req
}

func (req *Request) Queries(v interface{}) *Request {
	if req.req == nil || req.err != nil {
		return req
	}

	return req
}
