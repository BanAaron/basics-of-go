package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"aaronbarratt.dev/go/femm/data"
)

func Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := request.URL.Query()["id"]
	if id != nil {
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(writer).Encode(data.GetAll()[finalId])
		} else {
			http.Error(writer, "Invalid Exhibition", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(writer).Encode(data.GetAll())
	}
}
