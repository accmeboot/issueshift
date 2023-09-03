package response

import (
	"log"
	"net/http"

	"github.com/accmeboot/issueshift/internal/domain"
)

func WriteError(w http.ResponseWriter, statusCode int, errorDetails domain.Envelope, err error) {
	if err != nil {
		log.Println(err)
	}

	WriteJSON(w, statusCode, errorDetails)
}
