package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/inter-hubly/pilot/hlog"
)

func (e *environment) AddHealthEndpoint(ctx context.Context) {
	hlog.Info(ctx, "Server.AddHealthEndpoint", "Add Health endpoint")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"status":      "healthy",
			"environment": string(e.Config.Env),
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			return
		}
	})

	go http.ListenAndServe(fmt.Sprintf(":809%s", e.Config.Port[len(e.Config.Port)-1:]), nil)
}
