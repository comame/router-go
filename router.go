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

func (r *Router) All(pattern string, f handler) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, f)
	r.paramHandlers = append(r.paramHandlers, nil)
	r.methodFilters = append(r.methodFilters, "")
}

func (r *Router) AllDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, nil)
	r.paramHandlers = append(r.paramHandlers, f)
	r.methodFilters = append(r.methodFilters, "")
}

func (r *Router) Get(pattern string, f handler) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, f)
	r.paramHandlers = append(r.paramHandlers, nil)
	r.methodFilters = append(r.methodFilters, http.MethodGet)
}

func (r *Router) GetDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, nil)
	r.paramHandlers = append(r.paramHandlers, f)
	r.methodFilters = append(r.methodFilters, http.MethodGet)
}

func (r *Router) Post(pattern string, f handler) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, f)
	r.paramHandlers = append(r.paramHandlers, nil)
	r.methodFilters = append(r.methodFilters, http.MethodPost)
}

func (r *Router) PostDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, nil)
	r.paramHandlers = append(r.paramHandlers, f)
	r.methodFilters = append(r.methodFilters, http.MethodPost)
}

func (r *Router) Patch(pattern string, f handler) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, f)
	r.paramHandlers = append(r.paramHandlers, nil)
	r.methodFilters = append(r.methodFilters, http.MethodPatch)
}

func (r *Router) PatchDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, nil)
	r.paramHandlers = append(r.paramHandlers, f)
	r.methodFilters = append(r.methodFilters, http.MethodPatch)
}

func (r *Router) Delete(pattern string, f handler) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, f)
	r.paramHandlers = append(r.paramHandlers, nil)
	r.methodFilters = append(r.methodFilters, http.MethodDelete)
}

func (r *Router) DeleteDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, nil)
	r.paramHandlers = append(r.paramHandlers, f)
	r.methodFilters = append(r.methodFilters, http.MethodDelete)
}

func (r *Router) Put(pattern string, f handler) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, f)
	r.paramHandlers = append(r.paramHandlers, nil)
	r.methodFilters = append(r.methodFilters, http.MethodPut)
}

func (r *Router) PutDyn(pattern string, f handlerWithParam) {
	def := parseDef(pattern)
	r.rawPatterns = append(r.rawPatterns, pattern)
	r.defs = append(r.defs, def)
	r.handlers = append(r.handlers, nil)
	r.paramHandlers = append(r.paramHandlers, f)
	r.methodFilters = append(r.methodFilters, http.MethodPut)
}

func (router *Router) Route(w http.ResponseWriter, r *http.Request) {
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

func (router *Router) ListenAndServe(addr string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		router.Route(w, req)
	})
	return http.ListenAndServe(addr, nil)
}
