package methods

const (
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
)

var MethodList = [...]string{
	GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE,
}
