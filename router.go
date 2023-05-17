package router

import (
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)
type HandlerWithParam func(w http.ResponseWriter, r *http.Request, p Params)

type Params map[string]string

type Router struct {
	rawPatterns   []string
	defs          []routesDef
	handlers      []Handler
	paramHandlers []HandlerWithParam
	methodFilters []string
}

var router Router

func All(pattern string, f Handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, "")
}

func AllDyn(pattern string, f HandlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, "")
}

func Get(pattern string, f Handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, http.MethodGet)
}

func GetDyn(pattern string, f HandlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodGet)
}

func Post(pattern string, f Handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, http.MethodPost)
}

func PostDyn(pattern string, f HandlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodPost)
}

func Patch(pattern string, f Handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, http.MethodPatch)
}

func PatchDyn(pattern string, f HandlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodPatch)
}

func Delete(pattern string, f Handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, http.MethodDelete)
}

func DeleteDyn(pattern string, f HandlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodDelete)
}

func Put(pattern string, f Handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, http.MethodPut)
}

func PutDyn(pattern string, f HandlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodPut)
}

func Route(w http.ResponseWriter, r *http.Request) {
	pathstr := r.URL.Path
	for i, v := range router.defs {
		if matches(pathstr, v) && (router.methodFilters[i] == "" || r.Method == router.methodFilters[i]) {

			h := router.handlers[i]
			if h != nil {
				h(w, r)
				return
			}

			ph := router.paramHandlers[i]
			params := extractDynamicRoutes(pathstr, v)
			ph(w, r, params)
			return
		}
	}

	w.WriteHeader(404)
}

func ListenAndServe(addr string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		Route(w, req)
	})
	return http.ListenAndServe(addr, nil)
}
