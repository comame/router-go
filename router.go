package router

import (
	"context"
	"net/http"
)

type routerContextKey string

type Router struct {
	rawPatterns   []string
	defs          []routesDef
	handlers      []http.HandlerFunc
	methodFilters []string
}

var router Router

func All(pattern string, f http.HandlerFunc) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.methodFilters = append(router.methodFilters, "")
}

func Get(pattern string, f http.HandlerFunc) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodGet)
}

func Post(pattern string, f http.HandlerFunc) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodPost)
}

func Patch(pattern string, f http.HandlerFunc) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodPatch)
}

func Delete(pattern string, f http.HandlerFunc) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodDelete)
}

func Put(pattern string, f http.HandlerFunc) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodPut)
}

func Params(r *http.Request) map[string]string {
	pathstr := r.URL.Path
	def := r.Context().Value(routerContextKey("__router_def")).(routesDef)

	return extractDynamicRoutes(pathstr, def)
}

func Handler() http.Handler {
	return router
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route(w, r)
}

func route(w http.ResponseWriter, r *http.Request) {
	pathstr := r.URL.Path
	for i, def := range router.defs {
		if matches(pathstr, def) && (router.methodFilters[i] == "" || r.Method == router.methodFilters[i]) {
			h := router.handlers[i]
			k := routerContextKey("__router_def")
			ctx := context.WithValue(r.Context(), k, def)
			h(w, r.WithContext(ctx))
			return
		}
	}

	w.WriteHeader(404)
}
