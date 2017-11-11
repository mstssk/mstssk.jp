package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

func init() {
	http.HandleFunc("/api/health", handleAPIHealth)
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is dummy.")
}

func handleAPIHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := map[string]string{
		"runtime_version": runtime.Version(),
	}
	json, _ := json.Marshal(res)
	fmt.Fprint(w, string(json))
}
