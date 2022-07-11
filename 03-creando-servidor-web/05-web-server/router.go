package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func newRouter() *Router {
	return &Router{rules: make(map[string]map[string]http.HandlerFunc)}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handle, methodExist, exist := r.FindHandle(request.URL.Path, request.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusInternalServerError)
	}

	handle(w, request)

}

func (r *Router) FindHandle(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}
