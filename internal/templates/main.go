package templates

import (
	"bytes"
	"github.com/accmeboot/issueshift/view"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
)

type Cache map[string]*template.Template

func NewCache() (*Cache, error) {
	cache := Cache{}

	pages, err := fs.Glob(view.Files, "pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{
			"components/*.gohtml",
			page,
		}

		ts, err := template.New(name).ParseFS(view.Files, patterns...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return &cache, nil
}

func (c Cache) Render(w http.ResponseWriter, status int, page string, fragment *string, data any) {
	ts, ok := c[page]
	if !ok {
		log.Printf("template %s: does not exist", page)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// First executing a template into a buffer, in order to prevent corrupted templates to a client
	buf := new(bytes.Buffer)
	tmlFragment := "base"

	if fragment != nil {
		tmlFragment = *fragment
	}

	err := ts.ExecuteTemplate(buf, tmlFragment, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if _, err = buf.WriteTo(w); err != nil {
		log.Println(err)
	}
}

func (c Cache) ServerError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("HX-Redirect", "/error")
	w.WriteHeader(http.StatusInternalServerError)
}
