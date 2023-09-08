package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/accmeboot/issueshift/internal/domain"
	"io"
	"log"
	"net/http"
	"strings"
)

func (p *Provider) ReadBody(w http.ResponseWriter, r *http.Request, target any) error {
	maxBytes := 1_048_576 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(target)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown filed"):
			filedName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains unknown key %s", filedName)
		case errors.As(err, &invalidUnmarshalError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func (p *Provider) Send(w http.ResponseWriter, status int, data domain.Envelope) {
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

func (p *Provider) SendError(w http.ResponseWriter, statusCode int, errorDetails domain.Error, err error) {
	if err != nil {
		log.Println(err)
	}

	p.Send(w, statusCode, domain.Envelope{
		"errors": errorDetails,
	})
}

func (p *Provider) SendServerError(w http.ResponseWriter, err error) {
	p.SendError(w, http.StatusInternalServerError, domain.Error{
		"internal": "internal server error",
	}, err)
}

func (p *Provider) SendUnprocessableEntity(w http.ResponseWriter, err error) {
	p.SendError(w, http.StatusUnprocessableEntity, domain.Error{
		"unprocessable_entity": "failed to parse body",
	}, err)
}

func (p *Provider) SendBadRequest(w http.ResponseWriter, data domain.Error, err error) {
	p.SendError(w, http.StatusBadRequest, data, err)
}

func (p *Provider) SendNotFound(w http.ResponseWriter, err error) {
	p.SendError(w, http.StatusNotFound, domain.Error{
		"not_found": "requested resource could not be found",
	}, err)
}
