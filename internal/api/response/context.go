package response

import (
	"context"
	"net/http"
)

func PassToContext(r *http.Request, err error) {
	ctx := context.WithValue(r.Context(), ErrorKey, err)
	*r = *r.WithContext(ctx)
}
