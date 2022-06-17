package myrouter

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
)

type Middleware func(decoder *json.Decoder) (ok bool, status int, message string)

//RouterEntry is a struct that holds the method, path and handlerFunc
type RouteEntry struct {
	Path        *regexp.Regexp
	Method      string
	Middlewares []Middleware
	HandlerFunc http.HandlerFunc
}

//MyResponseWriter is a struct that holds the status code and the response writer
type MyResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (ent *RouteEntry) Match(r *http.Request) map[string]string {
	match := ent.Path.FindStringSubmatch(r.URL.Path)
	if match == nil {
		return nil // No match found
	}
	// Create a map to store URL parameters in
	params := make(map[string]string)
	groupNames := ent.Path.SubexpNames()
	for i, group := range match {
		params[groupNames[i]] = group
	}
	return params
}

//Router will send all requests to the router to be handled
type Router struct {
	routes []RouteEntry
}

//Route is a function that returns a router
func (rtr *Router) Route(method, path string, middlewares []Middleware, handlerFunc http.HandlerFunc) {
	e := RouteEntry{
		Method:      method,
		Path:        regexp.MustCompile(path),
		Middlewares: middlewares,
		HandlerFunc: handlerFunc,
	}
	rtr.routes = append(rtr.routes, e)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, e := range rtr.routes {
		params := e.Match(r)
		if params == nil {
			continue // No match found
		}
		var ok bool
		var status int
		var message string

		if e.Middlewares != nil {
			for _, middleware := range e.Middlewares {
				ok, status, message = middleware(decoder)
				if !ok {
					break
				}

			}
		}
		if ok {
			// Create new request with params stored in context
			ctx := context.WithValue(r.Context(), "params", params)
			e.HandlerFunc.ServeHTTP(w, r.WithContext(ctx))
		} else {

			defer r.Body.Close()
			w.WriteHeader(status)
			_, _ = w.Write([]byte("{\"message\":\"" + message + "\"}"))
		}

		return
	}
	http.NotFound(w, r)
}
