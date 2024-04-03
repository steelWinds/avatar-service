package app

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/steelWinds/identavatar/pkg"
)

func HandlerError(w http.ResponseWriter, externalErr error, status int) {
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(status)

	response := make(map[string]string)

	response["message"] = externalErr.Error()

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResponse)
}

func HandlerGetIdentcoin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, HEAD")
	w.Header().Set("Content-Type", "image/svg+xml")

	word := req.URL.Query().Get("word")

	if word == "" {
		HandlerError(w, errors.New("word is empty"), http.StatusBadRequest)

		return
	}

	squares, err := strconv.Atoi(req.URL.Query().Get("squares"))

	if err != nil {
		HandlerError(w, errors.New("squares is not int"), http.StatusBadRequest)

		return
	}

	size, err := strconv.Atoi(req.URL.Query().Get("size"))

	if err != nil {
		HandlerError(w, errors.New("size is not int"), http.StatusBadRequest)

		return
	}

	options := pkg.Options{
		Squares: squares,
		Size:    size,
		Word:    word,
	}

	buf, err := pkg.GetIndentcoin(options)

	if err != nil {
		HandlerError(w, errors.New("generate avatar failed"), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(buf.Bytes())
}
