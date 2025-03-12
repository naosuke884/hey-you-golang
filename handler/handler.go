package handler

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello, !!")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
