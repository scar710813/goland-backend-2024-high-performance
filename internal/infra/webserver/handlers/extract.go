package handlers

import "net/http"

func ExtractHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, extract!"))
}
