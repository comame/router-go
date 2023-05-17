package router_test

import (
	"fmt"
	"net/http"

	router "github.com/comame/router-go"
)

func Example() {
	router.GetDyn("/users/:userId", func(w http.ResponseWriter, r *http.Request, p router.Param) {
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
