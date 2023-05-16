Simple routing library for go.

## Usage

```go
// router_test.go

import (
	"fmt"
	"net/http"

	librouter "github.com/comame/router-go"
)

func Example() {
	router := new(librouter.Router)

	router.GetDyn("/users/:userId", func(w http.ResponseWriter, r *http.Request, p librouter.Param) {
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
```
