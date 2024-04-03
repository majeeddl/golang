package context1

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, t *http.Request) {
		// store.Cancel()
		// fmt.Fprintf(w, store.Fetch())

		// ctx := t.Context()

		// data := make(chan string, 1)

		// go func() {
		// 	data <- store.Fetch()
		// }()

		// select {
		// case d := <-data:
		// 	fmt.Fprintf(w, d)
		// case <-ctx.Done():
		// 	store.Cancel()
		// }

		data, _ := store.Fetch(r.Context())
		fmt.Fprint(w, data)
	}
}
