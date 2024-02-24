package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/steelWinds/identavatar/internal"
)

func main() {
	http.Handle("/", http.HandlerFunc(HandleGetIdentcoin))

	err := http.ListenAndServe(":3180", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func HandleError(w http.ResponseWriter, errorMessage string) {
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusInternalServerError)

	response := make(map[string]string)

	response["message"] = errorMessage

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResponse)
}

func HandleGetIdentcoin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, HEAD")
	w.Header().Set("Content-Type", "image/svg+xml")

	word := req.URL.Query().Get("word")

	squares, err := strconv.Atoi(req.URL.Query().Get("squares"))

	if err != nil {
		HandleError(w, "Squares is not int")

		return
	}

	size, err := strconv.Atoi(req.URL.Query().Get("size"))

	if err != nil {
		HandleError(w, "Size is not int")

		return
	}

	if word == "" || squares == 0 || size == 0 {
		HandleError(w, "All params is required")

		return
	}

	options := internal.Options{
		Squares: squares,
		Size:    size,
		Word:    word,
	}

	buf, err := internal.GetIndentcoin(options)

	if err != nil {
		HandleError(w, "Generate avatar failed")

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(buf.Bytes())
}
