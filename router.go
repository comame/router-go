package router

import (
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)
type handlerWithParam func(w http.ResponseWriter, r *http.Request, p Param)

type Param map[string]string

type Router struct {
	rawPatterns   []string
	defs          []routesDef
	handlers      []handler
	paramHandlers []handlerWithParam
	methodFilters []string
}

var router Router

func All(pattern string, f handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, "")
}

func AllDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, "")
}

func Get(pattern string, f handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, http.MethodGet)
}

func GetDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodGet)
}

func Post(pattern string, f handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, http.MethodPost)
}

func PostDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodPost)
}

func Patch(pattern string, f handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, http.MethodPatch)
}

func PatchDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodPatch)
}

func Delete(pattern string, f handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, http.MethodDelete)
}

func DeleteDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, nil)
	router.paramHandlers = append(router.paramHandlers, f)
	router.methodFilters = append(router.methodFilters, http.MethodDelete)
}

func Put(pattern string, f handler) {
	def := parseDef(pattern)
	router.rawPatterns = append(router.rawPatterns, pattern)
	router.defs = append(router.defs, def)
	router.handlers = append(router.handlers, f)
	router.paramHandlers = append(router.paramHandlers, nil)
	router.methodFilters = append(router.methodFilters, http.MethodPut)
}

func PutDyn(pattern string, f handlerWithParam) {
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
