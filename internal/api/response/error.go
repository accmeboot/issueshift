package response

import (
	"errors"
	"github.com/accmeboot/issueshift/internal/domain"
	"log"
	"net/http"
)

type contextKey string

var ErrorKey = contextKey("errorKey")

func ParseError(err error) (string, int) {
	switch {
	case errors.Is(err, domain.ErrNoRecord):
		return domain.ErrNoRecord.Error(), http.StatusNotFound
	case errors.Is(err, domain.ErrEditConflict):
		return domain.ErrEditConflict.Error(), http.StatusConflict
	default:
		log.Printf("internal server error: %s\n", err)
		return domain.ErrServer.Error(), http.StatusInternalServerError
	}
}
