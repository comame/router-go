package router_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	router "github.com/comame/router-go"
)

func Example() {
	router.Get("/users/:userId", func(w http.ResponseWriter, r *http.Request) {
		p := router.Params(r)
		fmt.Fprintln(w, "users/"+p["userId"])
	})

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	router.Post("/api/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from API")
	})

	router.ListenAndServe(":8080")
}

func ExampleGet() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	// matches /foo/bar
	router.Get("/foo/bar", handler)
	// matches /foo/bar, /foo/bar/baz, /foo/bar/baz/foo
	router.Get("/foo/bar/*", handler)
	router.Get("/*", handler)
}

func ExampleRoute() {
	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		router.Route(w, r)
	})
	http.ListenAndServe(":8080", nil)
}

func ExampleListenAndServe() {
	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!")
	})

	router.ListenAndServe(":8080")
}

func TestAll(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "https://exmaple.com/foo/bar", nil)
	w := httptest.NewRecorder()

	router.All("/foo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ng\n")
	})
	router.All("/foo/bar", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok\n")
	})
	router.All("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ng\n")
	})

	router.Route(w, r)

	if got := w.Body.String(); got != "ok\n" {
		t.Errorf("want %s, got %s\n", "ok", got)
	}
}

func TestGet(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "https://exmaple.com/foo/bar", nil)
	w := httptest.NewRecorder()

	router.Post("/foo/bar", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ngPost\n")
	})
	router.Get("/foo/bar", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok\n")
	})
	router.All("/foo/bar", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ngAll\n")
	})

	router.Route(w, r)

	if got := w.Body.String(); got != "ok\n" {
		t.Errorf("want %s, got %s\n", "ok", got)
	}
}

func TestParams(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "https://exmaple.com/users/1", nil)
	w := httptest.NewRecorder()

	router.All("/users/:user_id", func(w http.ResponseWriter, r *http.Request) {
		p := router.Params(r)
		io.WriteString(w, p["user_id"]+"\n")
	})

	router.Route(w, r)

	if got := w.Body.String(); got != "1\n" {
		t.Errorf("want %s, got %s\n", "1", got)
	}
}

func TestWildcard(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "https://exmaple.com/foo/bar", nil)
	w := httptest.NewRecorder()

	router.All("/foo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ng\n")
	})
	router.All("/foo/*", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok\n")
	})

	router.Route(w, r)

	if got := w.Body.String(); got != "ok\n" {
		t.Errorf("want %s, got %s\n", "ok", got)
	}
}
