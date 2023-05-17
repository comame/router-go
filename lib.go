package router

import (
	"strings"
)

type routeKind int

const (
	staticRoute routeKind = iota
	dynamicRoute
)

type routesDef struct {
	Routes           []routeDef
	TrailingWildcard bool
}

type routeDef struct {
	kind        routeKind
	staticName  string
	dynamicName string
}

func parseDef(pathstr string) routesDef {
	paths, trailingWildcard := splitPath(pathstr, false)

	routesDef := routesDef{
		Routes:           []routeDef{},
		TrailingWildcard: trailingWildcard,
	}

	for _, v := range paths {
		if v[0] == ':' {
			routesDef.Routes = append(routesDef.Routes, routeDef{
				kind:        dynamicRoute,
				dynamicName: v[1:],
			})
			continue
		}
		routesDef.Routes = append(routesDef.Routes, routeDef{
			kind:       staticRoute,
			staticName: v,
		})
	}

	return routesDef
}

func matches(pathstr string, def routesDef) bool {
	paths, _ := splitPath(pathstr, true)

	if !def.TrailingWildcard && (len(paths) != len(def.Routes)) {
		return false
	}

	if def.TrailingWildcard && (len(paths) < len(def.Routes)) {
		return false
	}

	for i, v := range def.Routes {
		if (v.kind == staticRoute) && (paths[i] != v.staticName) {
			return false
		}
		continue
	}

	return true
}

func extractDynamicRoutes(pathstr string, def routesDef) map[string]string {
	result := make(map[string]string)

	paths, _ := splitPath(pathstr, true)
	dynIndexes := make([]int, 0)

	for i, v := range def.Routes {
		if v.kind == dynamicRoute {
			dynIndexes = append(dynIndexes, i)
		}
	}

	for _, v := range dynIndexes {
		result[def.Routes[v].dynamicName] = paths[v]
	}

	return result
}

// 第 2 返値は末尾がワイルドカードか否か
func splitPath(pathstr string, includeTrailingWildcard bool) ([]string, bool) {
	pathsRaw := strings.Split(pathstr, "/")
	paths := make([]string, 0)

	trailingWildcard := false

	for i, v := range pathsRaw {
		if v == "" {
			continue
		}
		if i == len(pathsRaw)-1 && v == "*" {
			trailingWildcard = true
			if !includeTrailingWildcard {
				break
			}
		}
		paths = append(paths, v)
	}

	return paths, trailingWildcard
}
