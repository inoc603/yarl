package yarl

func (req *Request) Set(k string, v interface{}) *Request {
	req.err = req.body.SetItem(k, v)
	return req
}

func (req *Request) Body(v interface{}) *Request {
	req.err = req.body.Set(v)
	return req
}
