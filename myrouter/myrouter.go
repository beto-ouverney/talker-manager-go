package myrouter

import (
	"context"
	"net/http"
	"regexp"
)

//RouterEntry is a struct that holds the method, path and handlerFunc
type RouteEntry struct {
	Path        *regexp.Regexp
	Method      string
	HandlerFunc http.HandlerFunc
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

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	e := RouteEntry{
		Method:      method,
		Path:        regexp.MustCompile(path),
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
		// Create new request with params stored in context
		ctx := context.WithValue(r.Context(), "params", params)
		e.HandlerFunc.ServeHTTP(w, r.WithContext(ctx))
		return
	}
	http.NotFound(w, r)
}
