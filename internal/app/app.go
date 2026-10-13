package app

import (
	"fmt"
	"net/http"

	"github.com/kvakusha32/myapp/internal/app/handlers"
	"github.com/kvakusha32/myapp/utils"
)

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "Hello, Go project structure!")
	})

	mux.HandleFunc("/ping", handlers.Ping)

	mux.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		utils.WriteErr(w, http.StatusBadRequest, "bad_request_example")
	})

	utils.LogInfo("Server is starting on :8080")
	if err := http.ListenAndServe(":8080", withRequestID(mux)); err != nil {
		utils.LogError("server error: " + err.Error())
	}
}

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-Id")
		if id == "" {
			id = utils.NewID16()
		}
		r.Header.Set("X-Request-Id", id)
		w.Header().Set("X-Request-Id", id)
		next.ServeHTTP(w, r)
	})
}
