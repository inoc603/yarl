package yarl

const (
	POST    = "POST"
	GET     = "GET"
	HEAD    = "HEAD"
	PUT     = "PUT"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	OPTIONS = "OPTIONS"
)

func Get(url string) *Request {
	return newReq(GET, url)
}

func Put(url string) *Request {
	return newReq(PUT, url)
}

func Post(url string) *Request {
	return newReq(POST, url)
}

func Delete(url string) *Request {
	return newReq(DELETE, url)
}

func Patch(url string) *Request {
	return newReq(PATCH, url)
}
