package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All is running"))
	})
	http.ListenAndServe(":80", nil)
}
