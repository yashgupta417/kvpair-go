package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/yashgupta417/kvpair-go/services"
)

type GetKeyResponse struct {
	Value string
}

func GetKey(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")

	value := services.GetKeyFromMap(key)
	response := GetKeyResponse{Value: value}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(jsonData))
}

func UpsertKey(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	value := query.Get("value")

	services.UpsertKeyToMap(key, value)
	w.WriteHeader(200)
}

func DeleteKey(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	
	services.DeleteKeyFromMap(key)
	w.WriteHeader(200)
}