package gorctx

import (
	"context"
	"net/http"
)

// Set assigns value v under key k on given Request r's context.
func Set(r *http.Request, k, v interface{}) {
	if v == nil {
		return
	}
	*r = *r.WithContext(context.WithValue(r.Context(), k, v))
}

// Get retrieves value registered under key k of given Request context.
func Get(r *http.Request, k interface{}) interface{} {
	return r.Context().Value(k)
}

// GetOk retrieves value of key k from the given Request and indicates success or
// failure in 2nd return value.
func GetOk(r *http.Request, k interface{}) (v interface{}, ok bool) {
	if v = Get(r, k); v != nil {
		ok = true
	}
	return
}
