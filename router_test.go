package router_test

import (
	"fmt"
	"io"
	"net/http"

	router "github.com/comame/router-go"
)

func Example() {
	router.GetDyn("/users/:userId", func(w http.ResponseWriter, r *http.Request, p router.Params) {
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

func ExampleAll() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	// matches /foo/bar
	router.All("/foo/bar", handler)
	// matches /foo/bar, /foo/bar/baz, /foo/bar/baz/foo
	router.All("/foo/bar/*", handler)
	router.All("/*", handler)
}

func ExampleAllDyn() {
	handler := func(w http.ResponseWriter, r *http.Request, p router.Params) {
		userId := p["userId"]
		io.WriteString(w, "userId: "+userId)
	}

	router.AllDyn("/users/:userId", handler)
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

func ExampleGetDyn() {
	handler := func(w http.ResponseWriter, r *http.Request, p router.Params) {
		userId := p["userId"]
		io.WriteString(w, "userId: "+userId)
	}

	router.GetDyn("/users/:userId", handler)
}

func ExamplePost() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	// matches /foo/bar
	router.Post("/foo/bar", handler)
	// matches /foo/bar, /foo/bar/baz, /foo/bar/baz/foo
	router.Post("/foo/bar/*", handler)
	router.Post("/*", handler)
}

func ExamplePostDyn() {
	handler := func(w http.ResponseWriter, r *http.Request, p router.Params) {
		userId := p["userId"]
		io.WriteString(w, "userId: "+userId)
	}

	router.PostDyn("/users/:userId", handler)
}

func ExamplePatch() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	// matches /foo/bar
	router.Patch("/foo/bar", handler)
	// matches /foo/bar, /foo/bar/baz, /foo/bar/baz/foo
	router.Patch("/foo/bar/*", handler)
	router.Patch("/*", handler)
}

func ExamplePatchDyn() {
	handler := func(w http.ResponseWriter, r *http.Request, p router.Params) {
		userId := p["userId"]
		io.WriteString(w, "userId: "+userId)
	}

	router.PatchDyn("/users/:userId", handler)
}

func ExampleDelete() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	// matches /foo/bar
	router.Delete("/foo/bar", handler)
	// matches /foo/bar, /foo/bar/baz, /foo/bar/baz/foo
	router.Delete("/foo/bar/*", handler)
	router.Delete("/*", handler)
}

func ExampleDeleteDyn() {
	handler := func(w http.ResponseWriter, r *http.Request, p router.Params) {
		userId := p["userId"]
		io.WriteString(w, "userId: "+userId)
	}

	router.DeleteDyn("/users/:userId", handler)
}

func ExamplePut() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	// matches /foo/bar
	router.Put("/foo/bar", handler)
	// matches /foo/bar, /foo/bar/baz, /foo/bar/baz/foo
	router.Put("/foo/bar/*", handler)
	router.Put("/*", handler)
}

func ExamplePutDyn() {
	handler := func(w http.ResponseWriter, r *http.Request, p router.Params) {
		userId := p["userId"]
		io.WriteString(w, "userId: "+userId)
	}

	router.PutDyn("/users/:userId", handler)
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
