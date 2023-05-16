package router

import (
	"testing"
)

func TestMatches(t *testing.T) {
	def := parseDef("/foo/bar")
	result := matches("/foo/bar", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/foo/bar")
	result = matches("/foo/bar/", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/foo/bar")
	result = matches("/foo/bar/baz", def)
	if result {
		t.Errorf("")
	}

	def = parseDef("/foo/bar/*")
	result = matches("/foo/bar/baz", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/foo/bar/*")
	result = matches("/foo/bar/baz/foo", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/")
	result = matches("/", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/")
	result = matches("/foo", def)
	if result {
		t.Errorf("")
	}

	def = parseDef("/*")
	result = matches("/", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/*")
	result = matches("/foo", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/:id")
	result = matches("/", def)
	if result {
		t.Errorf("")
	}

	def = parseDef("/:id")
	result = matches("/foo", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/:id/*")
	result = matches("/foo", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/:id/*")
	result = matches("/foo/", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/:id/*")
	result = matches("/foo/bar", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/prefix/:id")
	result = matches("/prefix/foo", def)
	if !result {
		t.Errorf("")
	}

	def = parseDef("/prefix/:id")
	result = matches("/prefix", def)
	if result {
		t.Errorf("")
	}

	def = parseDef("/prefix/:id")
	result = matches("/foo", def)
	if result {
		t.Errorf("")
	}

	def = parseDef("/prefix/:id")
	result = matches("/foo/bar", def)
	if result {
		t.Errorf("")
	}
}

func TestExtractDynamicRoutes(t *testing.T) {
	def := parseDef("/:foo/:bar")
	m := extractDynamicRoutes("/foo/bar", def)

	if len(m) != 2 {
		t.Errorf("")
	}

	foo, ok := m["foo"]
	if !ok || foo != "foo" {
		t.Errorf("")
	}

	bar, ok := m["bar"]
	if !ok || bar != "bar" {
		t.Errorf("")
	}
}
