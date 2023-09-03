package response

import (
	"encoding/json"
	"github.com/accmeboot/issueshift/internal/domain"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data domain.Envelope) {
	info, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write(info); err != nil {
		log.Printf("error processing response json: %s\n", err)
	}
}
