// START 1 OMIT
// A Handler responds to an HTTP request.
//
// ...
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// END 1 OMIT

// START 2 OMIT
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

// END 2 OMIT