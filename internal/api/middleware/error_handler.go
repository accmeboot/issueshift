package middleware

import (
	"encoding/json"
	"github.com/accmeboot/issueshift/internal/api/response"
	"log"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			rawError := r.Context().Value(response.ErrorKey)
			if rawError == nil {
				return
			}

			errorData, ok := rawError.(error)
			if !ok {
				log.Printf("error accessing error details from context")
				http.Error(w, "internal server error", http.StatusInternalServerError)

				return
			}

			errorMessage, status := response.ParseError(errorData.(error))

			jsonData, err := json.Marshal(map[string]string{"error": errorMessage})
			if err != nil {
				log.Printf("failed to format errors: %s\n", err.Error())
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			if _, err = w.Write(jsonData); err != nil {
				log.Printf("failed to write errors: %s\n", err.Error())
			}
		}()
		next.ServeHTTP(w, r)
	})
}
