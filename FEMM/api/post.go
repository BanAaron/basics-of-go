package api

import (
	"encoding/json"
	"net/http"

	"aaronbarratt.dev/go/femm/data"
)

func Post(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var exhibition data.Exhibition
		err := json.NewDecoder(request.Body).Decode(&exhibition)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		data.Add(exhibition)
		writer.WriteHeader(http.StatusCreated)
		writer.Write([]byte("OK"))
	} else {
		http.Error(writer, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}
