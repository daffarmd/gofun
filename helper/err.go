package helper

import "net/http"

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HttpError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
}
