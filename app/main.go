package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/steelWinds/identavatar/internal"
)

func main() {
	http.Handle("/", http.HandlerFunc(HandleGetIdentcoin))

	err := http.ListenAndServe(":3180", nil)

	if err != nil {
		log.Fatal("Server shutdown:", err)
	}
}

func HandleError(w http.ResponseWriter, externalErr error) {
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusInternalServerError)

	response := make(map[string]string)

	response["message"] = externalErr.Error()

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

	if word == "" {
		HandleError(w, errors.New("word is empty"))

		return
	}

	squares, err := strconv.Atoi(req.URL.Query().Get("squares"))

	if err != nil {
		HandleError(w, errors.New("squares is not int"))

		return
	}

	size, err := strconv.Atoi(req.URL.Query().Get("size"))

	if err != nil {
		HandleError(w, errors.New("size is not int"))

		return
	}

	options := internal.Options{
		Squares: squares,
		Size:    size,
		Word:    word,
	}

	buf, err := internal.GetIndentcoin(options)

	if err != nil {
		HandleError(w, errors.New("generate avatar failed"))

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(buf.Bytes())
}
