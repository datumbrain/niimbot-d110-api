package main

import (
	"encoding/json"
	"net/http"
)

func printHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		setStatus(w, http.StatusBadRequest, "bad request")
		return
	}

	err = req.Validate()
	if err != nil {
		setStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	err = PrintTag(req.Text, req.QrText)
	if err != nil {
		setStatus(w, http.StatusInternalServerError, err.Error())
		return
	}

	setStatus(w, http.StatusOK, "tag printed")
}

func setStatus(w http.ResponseWriter, code int, msg string) error {
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(&Response{Status: msg})
}

func HandleAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUsername, reqPassword, ok := r.BasicAuth()
		if !ok || reqUsername != "DatumbrainHQ" || reqPassword != "password123" {
			setStatus(w, http.StatusUnauthorized, "Wrong Credentials")
			return
		}
		next.ServeHTTP(w, r)
	})
}
