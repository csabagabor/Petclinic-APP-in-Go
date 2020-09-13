package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Hello struct {
	Message string
	Error   int
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello There")
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		vals, ok := r.URL.Query()["site"]
		if !ok || len(vals[0]) < 1 {
			fmt.Fprint(w, "Error")
			return
		}

		resp, _ := http.Get(vals[0])
		fmt.Fprint(w, resp.Status)
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		marshal, _ := json.Marshal(Hello{"Not Found", 404})
		w.Write(marshal)
	})

	http.ListenAndServe(":8080", nil)
}
